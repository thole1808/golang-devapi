package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	// Post Data
	router.POST("/post", postHandler)

	// run port
	router.Run(":8282")
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

// POST DATA & JSON Validation & Show Error
type Post struct {
	Title string      `json:"title" binding:"required"`
	Harga interface{} `json:"harga" binding:"required,number"`
}

func postHandler(c *gin.Context) {
	var postInput Post
	err := c.ShouldBindJSON(&postInput)

	if err != nil {

		// 	// jika pakai log ini yang ini akan exit secara automatic
		// 	// log.Fatal(err)

		// 	// handling & show error json
		// 	// c.JSON(http.StatusBadRequest, err)
		// 	// fmt.Println(err)
		// 	// return

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, conditon: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": postInput.Title,
		"harga": postInput.Harga,
	})

}
