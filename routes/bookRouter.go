package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mcharolabs/go-crud/controllers"
)

func BookRouter(incomingRoutes *gin.Engine) {
	bookRoutes := incomingRoutes.Group("/books")
	{
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
		bookRoutes.PATCH("/:id", controllers.UpdateBook)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.POST("/", controllers.CreateBook)
		bookRoutes.GET("/", controllers.GetAllBooks)
	}
}
