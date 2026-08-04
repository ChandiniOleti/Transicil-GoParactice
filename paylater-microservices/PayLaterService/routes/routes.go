package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// All domain APIs have been extracted:
	// Auth :8081, User :8082, Merchant :8083,
	// Transaction :8084, Payback :8085, Report :8086.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "PayLater monolith (reports moved to :8086)",
		})
	})
}
