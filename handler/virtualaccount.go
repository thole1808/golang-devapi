package handler

import (
	"fmt"
	"golang-devapi/virtualaccount"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// root Index
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "null",
	})
}

// Get ID
func GetIdHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Get ID Multiple Path Params
func GetMultipleIdHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

// Query Handler
func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	judul := c.Query("judul")
	c.JSON(http.StatusOK, gin.H{"title": title, "judul": judul})
}

func PostHandler(c *gin.Context) {
	var postInput virtualaccount.Post
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
