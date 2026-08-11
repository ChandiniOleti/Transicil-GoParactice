package routes

import (
	"github.com/gin-gonic/gin"

	"paybackservice/config"
	"paybackservice/handlers"
	"paybackservice/middleware"

	"paylater.dev/shared/server"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", server.HealthHandler(config.ServiceName, nil))

	router.POST("/payback",
		middleware.JWTMiddleware(),
		handlers.Payback,
	)
}
