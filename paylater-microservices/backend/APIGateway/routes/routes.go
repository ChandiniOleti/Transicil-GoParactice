package routes

import (
	"github.com/gin-gonic/gin"

	"apigateway/config"
	"apigateway/middleware"
	"apigateway/proxy"

	"paylater.dev/shared/server"
)

func SetupRoutes(router *gin.Engine) {
	auth := proxy.Forward(config.AuthServiceURL)
	user := proxy.Forward(config.UserServiceURL)
	merchant := proxy.Forward(config.MerchantServiceURL)
	transaction := proxy.Forward(config.TransactionServiceURL)
	payback := proxy.Forward(config.PaybackServiceURL)
	report := proxy.Forward(config.ReportServiceURL)

	jwt := middleware.JWTMiddleware()

	router.GET("/health", server.HealthHandler(config.ServiceName, nil))

	// AUTH — no gateway rate limiting (public limits applied at AuthService)
	router.GET("/login/public-key", auth)
	router.POST("/login", auth)
	router.POST("/admin/login", auth)
	router.POST("/merchant/login", auth)
	router.POST("/admins", jwt, auth)
	router.GET("/admins", jwt, auth)

	// USERS
	router.POST("/users", user)
	router.GET("/users", jwt, user)
	router.GET("/users/:id", jwt, user)
	router.PUT("/users/:id", jwt, user)
	router.DELETE("/users/:id", jwt, user)

	// MERCHANTS
	router.POST("/merchants", merchant)
	router.GET("/merchants", jwt, merchant)
	router.GET("/merchants/:id", jwt, merchant)
	router.PUT("/merchants/:id", jwt, merchant)
	router.PATCH("/merchants/:id/commission", jwt, merchant)
	router.DELETE("/merchants/:id", jwt, merchant)

	// TRANSACTIONS (specific paths before /:id)
	router.POST("/transactions", jwt, transaction)
	router.GET("/transactions", jwt, transaction)
	router.GET("/transactions/user/:user_id", jwt, transaction)
	router.GET("/transactions/merchant/:merchant_id", jwt, transaction)
	router.GET("/transactions/:id", jwt, transaction)

	// PAYBACK
	router.POST("/payback", jwt, payback)

	// REPORTS
	router.GET("/reports/merchant-fee/:id", jwt, report)
	router.GET("/reports/users-with-due", jwt, report)
	router.GET("/reports/user-due/:id", jwt, report)
	router.GET("/reports/credit-limit-users", jwt, report)
	router.GET("/reports/total-dues", jwt, report)
}
