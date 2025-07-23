package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mcharolabs/go-crud/database"
	"github.com/mcharolabs/go-crud/helpers"
	"github.com/mcharolabs/go-crud/models"
	"github.com/mcharolabs/go-crud/utils/logger"
)

func GetUsers(c *gin.Context) {

}

func GetUser(c *gin.Context) {
	userID := c.Param("id")

	if err := helpers.MatchUserTypeToUUID(c, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbCtx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User

	result := database.DB.WithContext(dbCtx).First(&user, "id = ?", userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		logger.Error("User not found: " + result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
