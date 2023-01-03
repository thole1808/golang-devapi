package main

import (
	"golang-devapi/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Database *gorm.DB

// tester
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

	// Konfigurasi Database
	// var err error
	// host := os.Getenv("DB_HOST")
	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// databaseName := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	// Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Successfully connected to the database")
	// }
	// Konfigurasi Database

	// run port
	router.Run(":8282")
}
