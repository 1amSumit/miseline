package main

import (
	"github.com/gin-gonic/gin"
	"sumit.com/mise-link/db"
	"sumit.com/mise-link/routes"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080")
}
