package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mcharolabs/go-crud/controllers"
)

func AuthRoute(incoming *gin.Engine) {
	auth := incoming.Group("/auth")
	{
		auth.POST("/", controllers.SignIn)
		auth.POST("/sign-up", controllers.SignUp)
	}
}
