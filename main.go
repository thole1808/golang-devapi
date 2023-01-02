package main

import (
	"golang-devapi/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// buatlah route anda masing - masing sesuai kebutuhan
	router.GET("/", handler.RootHandler)

	// Get ID single
	router.GET("/getid/:id", handler.GetIdHandler)

	// Get Multiple Path Query
	router.GET("/getid/:id/:title", handler.GetMultipleIdHandler)

	// Get Query
	router.GET("/query", handler.QueryHandler)

	// Post Data
	router.POST("/post", handler.PostHandler)

	// run port
	router.Run(":8282")
}
