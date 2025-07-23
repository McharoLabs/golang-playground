package services

import (
	"errors"
	"strings"

	"github.com/mcharolabs/go-crud/database"
	"github.com/mcharolabs/go-crud/helpers"
	"github.com/mcharolabs/go-crud/models"
)

func SignUp(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	result := database.DB.Create(user)
	if result.Error != nil {
		lowerErr := strings.ToLower(result.Error.Error())
		validationErrors := helpers.ValidationErrors{}

		if strings.Contains(lowerErr, "email") {
			validationErrors["email"] = "email already exists"
		}
		if strings.Contains(lowerErr, "phone") {
			validationErrors["phone_number"] = "phone number already exists"
		}

		if len(validationErrors) > 0 {
			return validationErrors
		}
		return result.Error
	}

	return nil
}

func SignIn(email, password string) (accessToken, refreshToken string, err error) {
	var user models.User
	result := database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return "", "", errors.New("invalid email or password")
	}

	if err := helpers.CheckPasswordHash(password, user.Password); err != nil {
		return "", "", errors.New("invalid email or password")
	}

	accessToken, refreshToken, err = helpers.GenerateTokenPair(helpers.SignedDetails{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Uid:       user.ID.String(),
		UserType:  user.UserType,
	})
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
