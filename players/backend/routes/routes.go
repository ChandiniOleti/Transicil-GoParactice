package routes

import (
	"github.com/gin-gonic/gin"

	"players/handlers"
)

func RegisterRoutes(
	router *gin.Engine,
	playerHandler *handlers.PlayerHandler,
) {
	router.GET("/players", playerHandler.GetAllPlayers)
	router.GET("/players/:id", playerHandler.GetPlayerByID)
}
