package api

import (
	"attendr/watcher/log"
	"attendr/watcher/sentinel"
	"time"

	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	log.Logger.Info("Starting API web server...")
	router := gin.Default()

	// /api/v1/asu/watcher
	router.GET("/api/v1/asu/watcher", getWatchedClasses)
	router.POST("/api/v1/asu/watcher", addWatchedClass)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Logger.Fatal("Failed to start web server: %v", err)
	}
}

func getWatchedClasses(c *gin.Context) {
	if sentinel.ClassQueue.IsLocked() {
		time.Sleep(1 * time.Second)
		log.Logger.Debug("Waiting for ClassQueue to unlock to return classList...")
	}

	classes := sentinel.GetWatchedClasses()
	c.JSON(200, classes)
}

func addWatchedClass(c *gin.Context) {
	classNbr := c.Query("classNbr")
	term := c.Query("term")

	toAdd := sentinel.ASUWC_Wrapper{
		ClassNumber: classNbr,
		Term:        term,
	}

	res, err := sentinel.AddClass(toAdd)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": err.Error(),
		})

		return
	}

	log.Logger.Debug("Adding class to ClassQueue:\n", "classNbr", classNbr, "term", term)
	c.JSON(200, gin.H{
		"Success": res,
	})
}
