package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized user type")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

func CheckUserType(c *gin.Context, userType string) (err error) {
	err = nil
	userType1 := c.GetString("user_type")
	if userType1 != userType {
		err = errors.New("Unauthorized user type")
		return err
	}
	return err
}
