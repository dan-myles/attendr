package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// class represents data about a class
type class struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	HasOpenSeats bool   `json:"hasOpenSeats"`
	Instructor   string `json:"price"`
	Number       string `json:"number"`
	Subject      string `json:"subject"`
	Code         string `json:"code"`
}

// struct literal to test the endpoints
var classes = []class{
	{
		ID:           "1",
		Title:        "Introduction to Go",
		HasOpenSeats: true,
		Instructor:   "John Doe",
		Number:       "101",
		Subject:      "CS",
		Code:         "99831",
	},

	{
		ID:           "2",
		Title:        "Intermediate Go",
		HasOpenSeats: true,
		Instructor:   "John Doe",
		Number:       "102",
		Subject:      "CS",
		Code:         "99832",
	},
}

func main() {
	router := gin.Default()

	// /api/v1/asu/watcher
	router.GET("/api/v1/asu/watcher", getWatchedClasses)
	router.POST("/api/v1/asu/watcher", postAlbums)

	// /api/v1/asu/watcher/:id
	router.GET("/api/v1/asu/watcher/:id", getWatchedClassByID)

	router.Run("localhost:8080")
}

func getWatchedClasses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, classes)
}

func postAlbums(c *gin.Context) {
	var newClass class

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newClass); err != nil {
		return
	}

	// Add the new class to the slice.
	classes = append(classes, newClass)
	c.IndentedJSON(http.StatusCreated, newClass)
}

// getWatchedClassByID locates the class whose ID value matches the id
func getWatchedClassByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an class whose ID value matches the parameter.
	for _, a := range classes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "class not found"})
}
