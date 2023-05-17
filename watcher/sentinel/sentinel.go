package sentinel

import (
	"attendr/watcher/ent"
	"attendr/watcher/utils"
	"attendr/watcher/utils/asu"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	Watchlist []ent.ASU_Watched_Class
	asuApiUrl string
)

func InitSentinel() {
	// Load asuApiUrl from .env
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Error("Error loading .env file")
	}
	asuApiUrl = os.Getenv("ASU_API_URL")
	if len(asuApiUrl) <= 0 {
		utils.Logger.Error("ASU_API_URL not set in .env file")
	}
}

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

// func makeRequest(class ASUWC_Wrapper) (ASUWC_Wrapper, error) {
// 	// Perform request
// 	searchUrl := fmt.Sprintf(
// 		"%s/classes?refine=Y&advanced=true&classNbr=%s&searchType=all&term=%s",
// 		apiURL,
// 		class.ClassNumber,
// 		class.Term)
//
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", searchUrl, nil)
// 	if err != nil {
// 		return class, err
// 	}
//
// 	req.Header = http.Header{
// 		"Authorization": {"Bearer null"},
// 	}
//
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return class, err
// 	}
//
// 	defer res.Body.Close()
//
// 	if res.StatusCode != http.StatusOK {
// 		// return class, fmt.Errorf("bad status: %s", res.Status)
//
// 		// USED TO TEST ON BAD API
// 		return class, nil
// 	}
//
// 	bodyBytes, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return class, err
// 	}
//
// 	var result ASU_API_Response
// 	if err := json.Unmarshal(bodyBytes, &result); err != nil {
// 		return class, err
// 	}
//
// 	// Check if there was a response
// 	if result.Total.Value == 0 {
// 		return class, errors.New("No classes found")
// 	}
//
// 	openSeats := result.Classes[0].SeatInfo.EnrlCap - result.Classes[0].SeatInfo.EnrlTot
//
// 	if openSeats <= 0 {
// 		return class, nil
// 	}
//
// 	return ASUWC_Wrapper{
// 		ClassNumber: class.ClassNumber,
// 		Term:        class.Term,
// 		OpenSeats:   openSeats,
// 	}, nil
// }
//
// /*
// * ASU Watched Class Wrapper
// * Since we don't need *all* the information a class has in the DB,
// * to be able to check if it open. We can pull the data we __do__ need.
// * This will save on memory and increased performance by 1.3x-1.5x
// * Internally we will work with Wrappers, however externally we work with
// * the ASU_API_Response struct.
// **/
// type ASUWC_Wrapper struct {
// 	ClassNumber string
// 	Term        string
// 	OpenSeats   int
// }
//
// // We use 2 queues to store the classes we want to watch.
// // The first is for batching go routines to check availability
// // The second queue that is declared later is for the sentinel
// // to check if a class is open. It will be a fixed size for faster
// // read and writes from a concurrent queue.
// var (
// 	ClassQueue          = goconcurrentqueue.NewFIFO()
// 	scraping       bool = false
// 	apiURL         string
// 	externalApiURL string
// 	rateLimit      int
// 	watchInterval  int
// )
//
// // Exported Function to add a class to the queue
// func AddClass(class ASUWC_Wrapper) (string, error) {
// 	for scraping {
// 		time.Sleep(1 * time.Second)
// 		log.Logger.Debug("Waiting for scraping to finish before adding class to queue...")
// 	}
//
// 	// check if class exists
// 	classes := GetWatchedClasses()
// 	for _, c := range classes {
// 		if c.ClassNumber == class.ClassNumber {
// 			return "Class is already in queue, added user to watchers", nil
// 		}
// 	}
//
// 	err := ClassQueue.Enqueue(class)
// 	if err != nil {
// 		return "Error adding class to queue!", errors.New("Error adding class to queue")
// 	}
//
// 	log.Logger.Debug("Added class to ClassQueue:\n", "class", class)
//
// 	return "Class added to queue, added user to watchers", nil
// }
//
// // Exported Function to start the sentinel
// func Start(wg *sync.WaitGroup) {
// 	defer wg.Done()
//
// 	AddClass(ASUWC_Wrapper{ClassNumber: "88655", Term: "2237", OpenSeats: 0})
//
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Logger.Fatal("Error loading .env file")
// 	}
//
// 	rateLimit, _ = strconv.Atoi(os.Getenv("RATE_LIMIT"))
// 	watchInterval, _ = strconv.Atoi(os.Getenv("WATCH_INTERVAL"))
// 	apiURL = os.Getenv("TEST_API_URL")
//
// 	for {
// 		log.Logger.Info("Starting sentinel...")
//
// 		for ClassQueue.IsLocked() {
// 			log.Logger.Debug("ClassQueue is locked, waiting...")
// 			time.Sleep(1 * time.Second)
// 		}
//
// 		scraping = true
// 		startWatching()
// 		scraping = false
//
// 		log.Logger.Info("Checking again after WATCH_INTERVAL minutes...",
// 			"WATCH_INTERVAL", watchInterval)
// 		time.Sleep(time.Duration(watchInterval) * time.Second)
// 	}
// }
//
// // Internal Function to start watching classes
// func startWatching() {
// 	log.Logger.Debug("Starting to scan classes...")
// 	if ClassQueue.GetLen() <= 0 {
// 		log.Logger.Error("No classes found in ClassQueue")
// 		return
// 	}
//
// 	log.Logger.Debug("Class Queue Length:", "len", ClassQueue.GetLen())
//
// 	sentinelQueue := goconcurrentqueue.NewFixedFIFO(ClassQueue.GetLen())
//
// 	// Wait until it is unlocked
// 	for ClassQueue.IsLocked() {
// 		log.Logger.Debug("ClassQueue is locked while transferring, waiting...")
// 		time.Sleep(1 * time.Second)
// 	}
//
// 	// Move all classes from ClassQueue to sentinelQueue
// 	classQueueLen := ClassQueue.GetLen()
// 	for i := 0; i < classQueueLen; i++ {
// 		class, _ := ClassQueue.Dequeue()
// 		next := class.(ASUWC_Wrapper)
// 		err := sentinelQueue.Enqueue(next)
// 		if err != nil {
// 			log.Logger.Error("Error adding class to sentinelQueue")
// 		}
//
// 		log.Logger.Debug("Added class to sentinelQueue:\n", "class", next)
// 	}
//
// 	log.Logger.Debug("Finished adding classes to sentinelQueue")
// 	log.Logger.Debug("Sentinel Queue Length:", "len", sentinelQueue.GetLen())
//
// 	start := time.Now()
// 	defer func() {
// 		log.Logger.Info("Finished checking classes in:",
// 			"time", time.Since(start))
// 	}()
//
// 	if apiURL == "" {
// 		log.Logger.Fatal("{SCHOOL_CODE}_API_URL is blank!")
// 		return
// 	}
//
// 	wg := sync.WaitGroup{}
// 	ch := make(chan ASUWC_Wrapper)
// 	senQueueLen := sentinelQueue.GetLen()
//
// 	for i := 0; i < senQueueLen; i++ {
// 		log.Logger.Debug("Starting go routine:", "worker", i+1)
// 		wg.Add(1)
// 		cl, _ := sentinelQueue.Dequeue()
// 		next := cl.(ASUWC_Wrapper)
//
// 		go func(class ASUWC_Wrapper) {
// 			defer wg.Done()
// 			result, err := makeRequest(class)
// 			if err != nil {
// 				log.Logger.Error("Error making request to ASU API for this class",
// 					"err", err,
// 					"class", class)
// 				return
// 			} else if result.OpenSeats > 0 {
// 				foundOpenClass(result)
// 			} else {
// 				foundClosedClass(result)
// 			}
//
// 			ch <- result
// 		}(next)
// 		time.Sleep(time.Duration(rateLimit) * time.Millisecond)
// 	}
//
// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()
//
// 	for class := range ch {
// 		log.Logger.Debug("Flushing the following from the channel:", "class", class)
// 	}
// }
//
//
// // WORK NEEDS TO BE DONE HERE TO HANDLE THESE CASES
// func foundOpenClass(class ASUWC_Wrapper) {
// 	log.Logger.Info("Found open class!", "class", class)
// }
//
// func foundClosedClass(class ASUWC_Wrapper) {
// 	log.Logger.Debug("Class has no open seats", "class", class)
// 	log.Logger.Debug("Adding class back to ClassQueue", "class", class)
//
// 	err := ClassQueue.Enqueue(class)
// 	if err != nil {
// 		log.Logger.Error("Error adding class back to ClassQueue", "class", class,
// 			"error", err)
// 	}
// }
//
//
// func GetWatchedClasses() []ASUWC_Wrapper {
// 	for scraping {
// 		log.Logger.Debug("Scraping in progress, waiting...")
// 		time.Sleep(1 * time.Second)
// 	}
//
// 	length := ClassQueue.GetLen()
// 	var classes []ASUWC_Wrapper
//
// 	for i := 0; i < length; i++ {
// 		temp, err := ClassQueue.Get(i)
// 		if err != nil {
// 			log.Logger.Error("Error getting class from queue", "error", err)
// 		}
// 		classes = append(classes, temp.(ASUWC_Wrapper))
// 	}
//
// 	log.Logger.Debug("Returning watched classes to API function...")
// 	return classes
// }
