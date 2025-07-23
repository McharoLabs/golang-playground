package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mcharolabs/go-crud/models"
	"github.com/mcharolabs/go-crud/response"
	"github.com/mcharolabs/go-crud/services"
	"github.com/mcharolabs/go-crud/utils/logger"
)

// SignUp handler
func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response.Send[any](c.Writer, http.StatusBadRequest, false, "Invalid request body", nil, c.FullPath())
		return
	}

	// Validate with detailed errors
	if err := user.Validate(); err != nil {
		validationErrors := make(map[string]string)
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				validationErrors[e.Field()] = e.Tag()
			}
		} else {
			validationErrors["error"] = err.Error()
		}
		response.Send(c.Writer, http.StatusBadRequest, false, "Validation failed", validationErrors, c.FullPath())
		return
	}

	// Call service
	if err := services.SignUp(&user); err != nil {
		errorResponse := make(map[string]string)
		errMsg := err.Error()

		if strings.Contains(errMsg, "email") {
			errorResponse["email"] = "email already exists"
		}
		if strings.Contains(errMsg, "phone") {
			errorResponse["phone_number"] = "phone number already exists"
		}

		// If we found at least one duplicate error, respond with 409 and all errors
		if len(errorResponse) > 0 {
			response.Send(c.Writer, http.StatusConflict, false, "Duplicate field error", errorResponse, c.FullPath())
			return
		}

		// Otherwise generic error
		logger.Error("SignUp: " + err.Error())
		response.Send[any](c.Writer, http.StatusInternalServerError, false, "Failed to create new user", nil, c.FullPath())
		return
	}

	// Sanitize user response (exclude password)
	userResponse := map[string]interface{}{
		"id":           user.ID,
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"user_type":    user.UserType,
		"created_at":   user.CreatedAt,
		"updated_at":   user.UpdatedAt,
	}

	logger.Info("SignUp: user " + user.Email + " created successfully")
	response.Send(c.Writer, http.StatusCreated, true, "User registered successfully", userResponse, c.FullPath())
}

// SignIn handler
func SignIn(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Send[any](c.Writer, http.StatusBadRequest, false, "Invalid login credentials format", nil, c.FullPath())
		return
	}

	accessToken, refreshToken, err := services.SignIn(input.Email, input.Password)
	if err != nil {
		response.Send[any](c.Writer, http.StatusUnauthorized, false, "Invalid email or password", nil, c.FullPath())
		return
	}

	data := map[string]string{
		"access":  accessToken,
		"refresh": refreshToken,
	}

	response.Send(c.Writer, http.StatusOK, true, "Login successful", data, c.FullPath())
}
