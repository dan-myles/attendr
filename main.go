package main

import (
	"attendr/watcher/api"
	"attendr/watcher/database"
	"attendr/watcher/log"
	"attendr/watcher/sentinel"
	"sync"
)

func main() {
	// Setting Logging Preferences
	log.SetLogOptions()

	// Run schema migration tool
	database.Init()

	// Instantiate WaitGroup
	var wg sync.WaitGroup

	// Start the sentinel before the API
	wg.Add(1)
	go sentinel.Start(&wg)

	// Start the webserver
	api.StartWebServer()

	// Keep main thread alive
	wg.Wait()
}
