package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"players/config"
	"players/db/generated"
	"players/handlers"
	"players/routes"
	"players/services"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	queries := generated.New(db)

	playerService := services.NewPlayerService(queries)
	playerHandler := handlers.NewPlayerHandler(playerService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: false,
	}))

	routes.RegisterRoutes(router, playerHandler)

	log.Println("Server running on :8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}




//main.go
//    ↓
// Database connection
//    ↓
// SQLC Queries
//    ↓
// PlayerService
//    ↓
// PlayerHandler
//    ↓
// RegisterRoutes()
//    ↓
// Gin Server :8080