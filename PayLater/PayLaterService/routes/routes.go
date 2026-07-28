package routes

import (
	"github.com/gin-gonic/gin"

	"paylaterservice/handlers"
)

func SetupRoutes(router *gin.Engine) {

	// User APIs
	router.POST("/users", handlers.CreateUser)
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUserByID)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)


	// Merchant Routes
	router.POST("/merchants", handlers.CreateMerchant)
	router.GET("/merchants", handlers.GetMerchants)
	router.GET("/merchants/:id", handlers.GetMerchantByID)
	router.PUT("/merchants/:id", handlers.UpdateMerchant)
	router.PATCH("/merchants/:id/commission", handlers.UpdateCommission)
	router.DELETE("/merchants/:id", handlers.DeleteMerchant)

}