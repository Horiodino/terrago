package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckuserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")

	if userType != role {
		err := errors.New("Unauthorized")
		return err
	}

	return err
}

func MatchUser(c *gin.Context, Userid string) (err error) {
	usertype := c.GetString("user_type")
	uid := c.GetString("uid")

	if usertype == "USER" && uid != Userid {
		err = errors.New("Unauthorized")
	}

	err = CheckuserType(c, usertype)
	return err
}
