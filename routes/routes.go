package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("api/user/signup", createUser)
	server.POST("api/user/login", login)

	server.POST("/api/outlet", createOutlet)
}
