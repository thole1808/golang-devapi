package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// buatlah route anda masing - masing sesuai kebutuhan
	router.GET("/", rootHandler)

	// Get ID single
	router.GET("/getid/:id", getIdHandler)

	// Get Multiple Path Query
	router.GET("/getid/:id/:title", getMultipleIdHandler)

	// Get Query
	router.GET("/query", queryHandler)

	router.Run(":8181")
}

// root Index
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "null",
	})
}

// Get ID
func getIdHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Get ID Multiple Path Params
func getMultipleIdHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

// Query Handler
func queryHandler(c *gin.Context) {
	title := c.Query("title")
	judul := c.Query("judul")
	c.JSON(http.StatusOK, gin.H{"title": title, "judul": judul})
}
