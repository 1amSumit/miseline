package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sumit.com/mise-link/utils"
)

func IsLoggedIn(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "you are not logged in. Please log in to continue.",
		})
	}
	user_id, _, err := utils.VerifyJwtToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "you are not logged in. Please log in to continue.",
		})
	}

	c.Set("user_id", user_id)

	c.Next()
}
