package utils

import (
	"github.com/gin-gonic/gin"
)

func IsLoggedIn(context *gin.Context) bool {
	token := context.Request.Header.Get("Authorization")

	if token == "" {

		return false
	}
	err := verifyJwtToken(token)
	return err == nil
}
