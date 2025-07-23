package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	UserType  string
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateTokenPair(details SignedDetails) (accessToken string, refreshToken string, err error) {
	// Clone the claims
	accessClaims := SignedDetails{
		Email:     details.Email,
		FirstName: details.FirstName,
		LastName:  details.LastName,
		Uid:       details.Uid,
		UserType:  details.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-crud",
			Subject:   details.Uid,
		},
	}

	refreshClaims := SignedDetails{
		Email:     details.Email,
		FirstName: details.FirstName,
		LastName:  details.LastName,
		Uid:       details.Uid,
		UserType:  details.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-crud",
			Subject:   details.Uid,
		},
	}

	// Create tokens
	accessJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// Sign tokens
	accessToken, err = accessJwt.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = refreshJwt.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
