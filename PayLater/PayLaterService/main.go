package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"paylaterservice/config"
	"paylaterservice/routes"
)

func main() {

	// Connect to MySQL
	err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database Connection Failed:", err)
	}

	// Initialize SQLC
	config.InitSQLC()

	// Create Gin Router
	router := gin.Default()

	// Register Routes
	routes.SetupRoutes(router)

	// Start Server
	router.Run(":8080")
}