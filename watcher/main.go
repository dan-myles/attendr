package main

import (
	"attendr/watcher/api"
	"attendr/watcher/database"
	"attendr/watcher/sentinel"
	"attendr/watcher/utils"
	"sync"
)

func main() {
	// Initialize the database
	database.InitDbConnection()

	// Initialize the logger
	utils.InitLogger()

	// Initialize the sentinel
	sentinel.InitSentinel()

	// Initialize the waitgroup
	var wg sync.WaitGroup

	// Start the sentinel
	wg.Add(1)
	go sentinel.StartSentinel(&wg)

	// Start the webserver
	api.StartWebServer()

	wg.Wait()
}
