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

	// Transaction Routes
	router.POST("/transactions", handlers.CreateTransaction)
	router.GET("/transactions", handlers.GetTransactions)
	router.GET("/transactions/:id", handlers.GetTransactionByID)
	router.GET("/transactions/user/:user_id", handlers.GetTransactionsByUser)
	router.GET("/transactions/merchant/:merchant_id", handlers.GetTransactionsByMerchant)

	//Payback Routes
	router.POST("/payback", handlers.Payback)

	// Reports Routes

	router.GET("/reports/merchant-fee/:id", handlers.MerchantFeeReport)
	router.GET("/reports/users-with-due", handlers.UsersWithDueReport)
	router.GET("/reports/user-due/:id", handlers.UserDueReport)
	router.GET("/reports/credit-limit-users", handlers.CreditLimitUsersReport)
	router.GET("/reports/total-dues", handlers.TotalDuesReport)

	//loginroutes
	router.POST("/login", handlers.LoginHandler)
}