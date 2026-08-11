package routes

import (
	"github.com/gin-gonic/gin"

	"reportservice/config"
	"reportservice/handlers"
	"reportservice/middleware"

	"paylater.dev/shared/server"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", server.HealthHandler(config.ServiceName, nil))

	router.GET("/reports/merchant-fee/:id",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.MerchantFeeReport,
	)

	router.GET("/reports/users-with-due",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.UsersWithDueReport,
	)

	router.GET("/reports/user-due/:id",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.UserDueReport,
	)

	router.GET("/reports/credit-limit-users",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.CreditLimitUsersReport,
	)

	router.GET("/reports/total-dues",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.TotalDuesReport,
	)
}
