// TODO:
// - Add email & phone number to DB
// - Add GET route to get watched classes for a certain user
// - Fix the DELETE route to:
//    First delete the users notification preferences
//    If it was the last user tracking the class then delete it from the DB
// Currently it just deletes the class from the DB and will affect all users
// - Add Clerk authentication to all routes

package api

import (
	"attendr/watcher/database"
	"attendr/watcher/sentinel"
	"attendr/watcher/utils"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	utils.Logger.Info("Starting API web server...")
	router := gin.Default()

	// /api/v1/asu/watcher
	router.GET("/api/v1/asu/watcher", getWatchedClasses)
	router.POST("/api/v1/asu/watcher", addWatchedClass)
	router.DELETE("/api/v1/asu/watcher", removeWatchedClass)

	err := router.Run("localhost:8080")
	if err != nil {
		utils.Logger.Fatal("Failed to start web server: %v", err)
	}
}

func addWatchedClass(c *gin.Context) {
	classNbr := c.Query("classNbr")
	term := c.Query("term")
	// email := c.Query("email")
	// phone := c.Query("phone")

	// First check if the class is already being watched
	for _, class := range sentinel.Watchlist {
		if class.ClassNumber == classNbr && class.Term == term {
			c.JSON(200, gin.H{
				"message": "Class already being watched.",
			})
			return
		}
	}

	// If it is not being watched, grab class info and watch it
	apiResponse, _ := sentinel.GrabClassInfo(term, classNbr)
	if apiResponse.Total.Value <= 0 {
		// The class code was invalid
		c.JSON(400, gin.H{
			"message": "No class found to add, try checking the term or class number.",
		})
		return
	}

	// Add class to DB
	class, err := database.AddClassToDb(context.Background(), &apiResponse)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error adding class to DB.",
		})
	}

	// Add class to sentinel.Watchlist
	// Have to dispatch go routine as we will not add a class while scraping
	go func() {
		for sentinel.Scraping {
			time.Sleep(1 * time.Second)
		}

		sentinel.Watchlist = append(sentinel.Watchlist, class)
		utils.Logger.Info("Added class to watchlist:", "class", class.ClassNumber)
	}()

	c.JSON(200, gin.H{
		"message": "Class in queue to be added to watchlist.",
	})
}

/*
* Removes a watched class
**/
func removeWatchedClass(c *gin.Context) {
	classNbr := c.Query("classNbr")
	term := c.Query("term")
	// email := c.Query("email")
	// phone := c.Query("phone")

	// First check if it is being watched
	// Check the DB
	class, err := database.GetClassFromDbByClassNbr(context.Background(), classNbr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error removing class from watchlist. It may not exist.",
		})
		return
	}

	// Remove class from DB
	err = database.RemoveClassFromDb(context.Background(), class)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error removing class from watchlist. It may not exist.",
		})
		return
	}

	// Remove class from setinel.Watchlist
	for i, class := range sentinel.Watchlist {
		if class.ClassNumber == classNbr && class.Term == term {
			sentinel.Watchlist = append(sentinel.Watchlist[:i], sentinel.Watchlist[i+1:]...)
			c.JSON(200, gin.H{
				"message": "Class removed from watchlist.",
			})
			return
		}
	}
}

/*
* Returns the array of watched classes
**/
func getWatchedClasses(c *gin.Context) {
	c.JSON(200, sentinel.Watchlist)
}
