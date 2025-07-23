package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mcharolabs/go-crud/controllers"
	"github.com/mcharolabs/go-crud/middleware"
)

func UserRoutes(incoimingRoutes *gin.Engine) {
	incoimingRoutes.Use(middleware.Authenticate())

	users := incoimingRoutes.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUser)
	}
}
