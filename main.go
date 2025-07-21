package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mcharolabs/go-crud/database"
	"github.com/mcharolabs/go-crud/routes"
	"github.com/mcharolabs/go-crud/utils/logger"
)

func init() {
	logger.Init()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.Connect()

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

}

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	routes.BookRouter(router)

	router.Run()
}
