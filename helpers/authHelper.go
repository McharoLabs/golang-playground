package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUUID(c *gin.Context, id string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userType {
		err = errors.New("unauthorized to access this resource")
		return err
	}

	err = checkUserType(c, userType)
	return err
}

func checkUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")

	err = nil
	if userType != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}

	return err
}
