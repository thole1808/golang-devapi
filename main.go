package main

import (
	"fmt"
	"golang-devapi/handler"
	"golang-devapi/virtualaccount"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

// tester
func main() {

	// Koneksi Ke Database PostgresSQL
	dsn := "host=localhost user=postgres password=1234 dbname=app-internal port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Koneksi ke DB")
	}
	fmt.Println("database sudah terkoneksi")

	// Repository
	kontrak := virtualaccount.NewRepository(db)
	kontraks, err := kontrak.FindAll()

	for _, virtualaccount := range kontraks {
		fmt.Println("test :", virtualaccount.No_invoice)
	}

	// ROUTE FRAMEWORKS GIN
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

	// Koneksikan Ke Database
	// db, err := model.Database()
	// if err != nil {
	//     log.Println(err)
	// }
	// db.DB()

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
