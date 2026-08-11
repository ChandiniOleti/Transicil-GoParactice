package routes

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"userservice/config"
	"userservice/handlers"
	"userservice/middleware"

	"paylater.dev/shared/ratelimit"
	"paylater.dev/shared/server"
)

func SetupRoutes(router *gin.Engine, limit rate.Limit, burst int) {
	router.GET("/health", server.HealthHandler(config.ServiceName, config.DB))

	publicLimit := ratelimit.Middleware(limit, burst)

	// Public registration (rate limited)
	router.POST("/users", publicLimit, handlers.CreateUser)

	// Admin-only list
	router.GET("/users",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.GetUsers,
	)

	// Own-account (or admin) user routes
	router.GET("/users/:id",
		middleware.JWTMiddleware(),
		middleware.UserMiddleware(),
		handlers.GetUserByID,
	)

	router.PUT("/users/:id",
		middleware.JWTMiddleware(),
		middleware.UserMiddleware(),
		handlers.UpdateUser,
	)

	router.DELETE("/users/:id",
		middleware.JWTMiddleware(),
		middleware.UserMiddleware(),
		handlers.DeleteUser,
	)

	// Internal service-to-service APIs (not rate limited)
	router.PATCH("/internal/users/:id/due",
		middleware.InternalMiddleware(),
		handlers.UpdateCurrentDue,
	)

	router.GET("/internal/users",
		middleware.InternalMiddleware(),
		handlers.GetUsersInternal,
	)

	router.GET("/internal/users/:id",
		middleware.InternalMiddleware(),
		handlers.GetUserByIDInternal,
	)
}
