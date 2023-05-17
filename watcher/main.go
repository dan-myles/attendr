package main

import (
	"attendr/watcher/api"
	"attendr/watcher/database"
	"attendr/watcher/sentinel"
	"attendr/watcher/utils"
)

func main() {
	// Initialize the database
	database.InitDbConnection()

	// Initialize the logger
	utils.InitLogger()

	// Initialize the sentinel
	sentinel.InitSentinel()

	// Start the API webserver
	api.StartWebServer()
}
