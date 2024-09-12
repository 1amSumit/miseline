package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/user/signup", createUser)
	server.POST("/user/login", login)
}
