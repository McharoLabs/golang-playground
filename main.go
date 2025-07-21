package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mcharolabs/go-crud/controllers"
	"github.com/mcharolabs/go-crud/database"
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
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	bookRoutes := router.Group("/books")
	{
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
		bookRoutes.PATCH("/:id", controllers.UpdateBook)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.POST("/", controllers.CreateBook)
		bookRoutes.GET("/", controllers.GetAllBooks)
	}

	router.Run()
}
