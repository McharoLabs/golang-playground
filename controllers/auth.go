package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mcharolabs/go-crud/database"
	"github.com/mcharolabs/go-crud/helpers"
	"github.com/mcharolabs/go-crud/models"
	"github.com/mcharolabs/go-crud/utils/logger"
)

func SignIn(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Validate request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login credentials format"})
		return
	}

	// Find user by email
	var user models.User
	result := database.DB.First(&user, "email = ?", input.Email)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	logger.Debug("Got user")

	if err := helpers.CheckPasswordHash(input.Password, user.Password); err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	accessToken, refreshToken, err := helpers.GenerateTokenPair(helpers.SignedDetails{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Uid:       user.ID.String(),
		UserType:  user.UserType,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	// Optionally store in DB
	user.Token = accessToken
	user.RefreshToken = refreshToken
	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"access":  accessToken,
		"refresh": refreshToken,
	})
}

func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := user.Validate(); err != nil {
		errs := err.(validator.ValidationErrors)
		validationErrors := make(map[string]string)

		for _, e := range errs {
			jsonTag := helpers.GetJSONTag(user, e.Field())
			validationErrors[jsonTag] = jsonTag + " is " + e.Tag()
		}

		c.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	// 🔐 Hash password before saving
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		logger.Error("SignUp: password hashing failed - " + err.Error())
		return
	}
	user.Password = hashedPassword

	// Save user
	result := database.DB.Create(&user)
	if result.Error != nil {
		errorResponse := make(map[string]string)
		lowerErr := strings.ToLower(result.Error.Error())

		if strings.Contains(lowerErr, "email") {
			errorResponse["email"] = "email already exists"
		}
		if strings.Contains(lowerErr, "phone") || strings.Contains(lowerErr, "phoneNumber") {
			errorResponse["phoneNumber"] = "phone number already exists"
		}

		if len(errorResponse) > 0 {
			c.JSON(http.StatusConflict, errorResponse)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new user"})
		logger.Error("CreateUser: failed to create user - " + result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	logger.Info("User with email: " + user.Email + " created successfully")
}
