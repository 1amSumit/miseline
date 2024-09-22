package utils

import (
	"github.com/gin-gonic/gin"
)

func IsLoggedIn(context *gin.Context) (float64, string, bool) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {

		return -1, "", false
	}
	user_id, email, err := VerifyJwtToken(token)
	if err != nil {
		return -1, "", false
	}

	return user_id, email, true
}
