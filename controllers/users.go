package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mcharolabs/go-crud/database"
	"github.com/mcharolabs/go-crud/models"
	"github.com/mcharolabs/go-crud/response"
	"github.com/mcharolabs/go-crud/utils/logger"
	"gorm.io/gorm"
)

// GetUsers returns all users (middleware should restrict access)
func GetUsers(c *gin.Context) {
	var users []models.User

	err := database.DB.Find(&users).Error
	if err != nil {
		logger.Error("GetUsers DB error: " + err.Error())
		response.Send[any](c.Writer, http.StatusInternalServerError, false, "Failed to fetch users", nil, c.FullPath())
		return
	}

	response.Send(c.Writer, http.StatusOK, true, "Users fetched successfully", users, c.FullPath())
}

// GetUser returns a single user by ID (middleware controls access)
func GetUser(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		response.Send[any](c.Writer, http.StatusBadRequest, false, "Invalid user ID", nil, c.FullPath())
		return
	}

	var user models.User
	err = database.DB.First(&user, "id = ?", userID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Send[any](c.Writer, http.StatusNotFound, false, "User not found", nil, c.FullPath())
			return
		}
		logger.Error("GetUser DB error: " + err.Error())
		response.Send[any](c.Writer, http.StatusInternalServerError, false, "Failed to fetch user", nil, c.FullPath())
		return
	}

	response.Send(c.Writer, http.StatusOK, true, "User fetched successfully", user, c.FullPath())
}
