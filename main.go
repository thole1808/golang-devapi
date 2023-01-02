package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// buatlah route anda masing - masing sesuai kebutuhan
	router.GET("/", rootHandler)
	router.GET("/getid/:id", getIdHandler)

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
