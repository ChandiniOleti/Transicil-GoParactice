package routes

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"merchantservice/config"
	"merchantservice/handlers"
	"merchantservice/middleware"

	"paylater.dev/shared/ratelimit"
	"paylater.dev/shared/server"
)

func SetupRoutes(router *gin.Engine, limit rate.Limit, burst int) {
	router.GET("/health", server.HealthHandler(config.ServiceName, config.DB))

	publicLimit := ratelimit.Middleware(limit, burst)

	// Public registration (rate limited)
	router.POST("/merchants", publicLimit, handlers.CreateMerchant)

	router.GET("/merchants",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.GetMerchants,
	)

	router.GET("/merchants/:id",
		middleware.JWTMiddleware(),
		middleware.MerchantMiddleware(),
		handlers.GetMerchantByID,
	)

	router.PUT("/merchants/:id",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.UpdateMerchant,
	)

	router.PATCH("/merchants/:id/commission",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.UpdateCommission,
	)

	router.DELETE("/merchants/:id",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.DeleteMerchant,
	)

	// Internal service-to-service APIs (not rate limited)
	router.GET("/internal/merchants/:id/commission",
		middleware.InternalMiddleware(),
		handlers.GetMerchantCommission,
	)
}
