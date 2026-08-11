package routes

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"authservice/config"
	"authservice/handlers"
	"authservice/middleware"

	"paylater.dev/shared/ratelimit"
	"paylater.dev/shared/server"
)

func SetupRoutes(router *gin.Engine, limit rate.Limit, burst int) {
	router.GET("/health", server.HealthHandler(config.ServiceName, config.DB))

	publicLimit := ratelimit.Middleware(limit, burst)

	// Public login routes (rate limited)
	router.POST("/login", publicLimit, handlers.LoginHandler)
	router.POST("/admin/login", publicLimit, handlers.AdminLoginHandler)
	router.POST("/merchant/login", publicLimit, handlers.MerchantLoginHandler)

	// Admin CRUD (JWT + Admin only)
	router.POST("/admins",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.CreateAdmin,
	)

	router.GET("/admins",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.GetAdmins,
	)
}
