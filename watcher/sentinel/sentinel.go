// TODO:
// - Handle finding open seats:
//    Seat is found ~ send notification & remove from DB + Watchlist
//    Seat is not found ~ do nothing and scrape it again later

package sentinel

import (
	"attendr/watcher/database"
	"attendr/watcher/ent"
	"attendr/watcher/utils"
	"attendr/watcher/utils/asu"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var (
	Watchlist       []*ent.ASUWatchedClass
	Scraping        bool   = false
	asuApiUrl       string = ""
	workerInterval  int    = 125 // Milliseconds
	scraperInterval int    = 240 // Minutes
)

/*
* Init sentinel by loading env variables, grabbing watchlist from db, etc
* If some env variables aren't set they have defaults
**/
func InitSentinel() {
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Error("Error loading .env file")
	}

	// Load asuApiUrl from .env
	asuApiUrl = os.Getenv("ASU_API_URL")
	if len(asuApiUrl) <= 0 {
		utils.Logger.Error("ASU_API_URL not set in .env file")
	}

	// Load worker interval
	temp := os.Getenv("WORKER_INTERVAL")
	if len(temp) <= 0 {
		utils.Logger.Error("WORKER_INTERVAL not set in .env file")
	}
	workerInterval, _ = strconv.Atoi(temp)

	// Load scraper interval
	temp2 := os.Getenv("SCRAPER_INTERVAL")
	if len(temp2) <= 0 {
		utils.Logger.Error("SCRAPER_INTERVAL not set in .env file")
	}
	scraperInterval, _ = strconv.Atoi(temp2)

	// Grab the most current watchlist from the db
	Watchlist, _ = database.GrabWatchListFromDb(context.Background())
}

/*
* Main function to hit ASU api and turn it into a struct asu.ApiResponse
**/
func GrabClassInfo(term string, classNbr string) (asu.ApiResponse, error) {
	// Load API URL from .env
	// We can do this separately to not worry about order of calling
	if len(asuApiUrl) <= 0 {
		err := godotenv.Load()
		if err != nil {
			utils.Logger.Error("Error loading .env file")
		}
		asuApiUrl = os.Getenv("ASU_API_URL")
	}

	// Format url
	searchUrl := fmt.Sprintf(
		"%s/classes?refine=Y&advanced=true&classNbr=%s&searchType=all&term=%s",
		asuApiUrl,
		classNbr,
		term)

	// Instantiate client
	client := &http.Client{}
	req, err := http.NewRequest("GET", searchUrl, nil)
	if err != nil {
		return asu.ApiResponse{}, err
	}

	// Set headers
	req.Header = http.Header{
		"Authorization": {"Bearer null"},
	}

	// Do request
	res, err := client.Do(req)
	if err != nil {
		return asu.ApiResponse{}, err
	}
	defer res.Body.Close()

	// Check status code
	if res.StatusCode != http.StatusOK {
		return asu.ApiResponse{}, fmt.Errorf("bad status: %s", res.Status)
	}

	// Convert response to bytes
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return asu.ApiResponse{}, err
	}

	// Unmarshal response
	var result asu.ApiResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return asu.ApiResponse{}, err
	}

	return result, nil
}

/*
* Loop forever to scrape the ASU API for class information
* Generally starts the sentinel
**/
func StartSentinel(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		utils.Logger.Info("Starting sentinel...")

		Scraping = true
		checkClasses()
		Scraping = false

		utils.Logger.Info("Sentinel sleeping for",
			"minutes",
			scraperInterval)
		time.Sleep(time.Duration(scraperInterval) * time.Minute)
	}
}

/*
* checkClasses will dispatch workers to grab class information
* from the ASU API. It will then check if the class is open or not.
**/
func checkClasses() {
	// Check if there are any classes to check
	if len(Watchlist) <= 0 {
		utils.Logger.Info("No classes found in Watchlist")
		return
	}

	// Track how long it takes to check all classes
	start := time.Now()
	defer func() {
		utils.Logger.Info("Finished checking classes in:",
			"time", time.Since(start))
	}()

	// Initialize the channel and waitgroup
	ch := make(chan asu.ApiResponse)
	wg := sync.WaitGroup{}

	// Loop through the Watchlist and dispatch workers
	for i := 0; i < len(Watchlist); i++ {
		wg.Add(1)
		utils.Logger.Debug("Starting go routine:", "worker", i+1)

		// Grab class & start worker
		classToScrape := Watchlist[i]
		go func(class *ent.ASUWatchedClass) {
			defer wg.Done()

			res, err := GrabClassInfo(class.Term, class.ClassNumber)
			if err != nil {
				utils.Logger.Error("Error grabbing class info:", "error", err)
			}

			handleResponse(res)

			ch <- res
		}(classToScrape)

		// Sleep for a bit to not overload the API
		utils.Logger.Debug("Waiting to dispatch next worker", "ms", workerInterval)
		time.Sleep(time.Duration(workerInterval) * time.Millisecond)
	}

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Flush the channel by using values in it
	for class := range ch {
		utils.Logger.Debug("Finished checking class:",
			"classNbr",
			class.Classes[0].Clas.Classnbr)
	}
}

/*
* Temporary function to handle open & closed classes
**/
func handleResponse(class asu.ApiResponse) {
	totalSeats, _ := strconv.Atoi(class.Classes[0].Clas.Enrlcap)
	availableSeats, _ := strconv.Atoi(class.Classes[0].Clas.Enrltot)

	openSeats := totalSeats - availableSeats

	if openSeats > 0 {
		utils.Logger.Info("Class has open seats:",
			"classNbr", class.Classes[0].Clas.Classnbr,
			"openSeats", openSeats)
	}

	if openSeats <= 0 {
		utils.Logger.Debug("Class has no open seats:",
			"classNbr", class.Classes[0].Clas.Classnbr)
	}
}
