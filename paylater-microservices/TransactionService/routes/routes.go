package routes

import (
	"github.com/gin-gonic/gin"

	"transactionservice/config"
	"transactionservice/handlers"
	"transactionservice/middleware"

	"paylater.dev/shared/server"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", server.HealthHandler(config.ServiceName, config.DB))

	router.POST("/transactions",
		middleware.JWTMiddleware(),
		handlers.CreateTransaction,
	)

	router.GET("/transactions",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.GetTransactions,
	)

	// Specific paths before /:id
	router.GET("/transactions/user/:user_id",
		middleware.JWTMiddleware(),
		middleware.UserMiddleware(),
		handlers.GetTransactionsByUser,
	)

	router.GET("/transactions/merchant/:merchant_id",
		middleware.JWTMiddleware(),
		middleware.MerchantMiddleware(),
		handlers.GetTransactionsByMerchant,
	)

	router.GET("/transactions/:id",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
		handlers.GetTransactionByID,
	)

	// Internal service-to-service: record PAYBACK rows
	router.POST("/internal/transactions/payback",
		middleware.InternalMiddleware(),
		handlers.CreatePaybackTransaction,
	)

	router.GET("/internal/transactions",
		middleware.InternalMiddleware(),
		handlers.GetTransactionsInternal,
	)

	router.GET("/internal/transactions/user/:user_id",
		middleware.InternalMiddleware(),
		handlers.GetTransactionsByUserInternal,
	)

	router.GET("/internal/transactions/merchant/:merchant_id",
		middleware.InternalMiddleware(),
		handlers.GetTransactionsByMerchantInternal,
	)
}
