package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"sumit.com/mise-link/db"
	"sumit.com/mise-link/routes"
)

func main() {

	db.InitDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
