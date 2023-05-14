package main

import (
	"attendr/watcher/log"
	"attendr/watcher/sentinel"
)

func main() {
	// Setting Logging Preferences
	log.SetLogOptions()

	sentinel.Start()
}
