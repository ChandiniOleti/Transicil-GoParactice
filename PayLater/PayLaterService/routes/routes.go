package routes

import (
	"github.com/gin-gonic/gin"

	"paylaterservice/handlers"
	"paylaterservice/middleware"
)

func SetupRoutes(router *gin.Engine) {

	// ==========================================================
	// PUBLIC ROUTES (No JWT Token Required)
	// ==========================================================
	// These APIs can be accessed without logging in.

	// User Registration
	router.POST("/users", handlers.CreateUser)

	// User Login
	router.POST("/login", handlers.LoginHandler)

	// Admin Login
	router.POST("/admin/login", handlers.AdminLoginHandler)

	//Merchant REgistration
	router.POST("/merchants",
		// middleware.JWTMiddleware(),
		// middleware.AdminMiddleware(),
		handlers.CreateMerchant,
	)

	// Merchant Login
	router.POST("/merchant/login", handlers.MerchantLoginHandler)

	// ==========================================================
	// USER ROUTES (JWT Token Required)
	// ==========================================================
	// These APIs work only after successful login.
	// Send:
	// Authorization: Bearer <JWT_TOKEN>

	router.GET("/users",
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(), // Uncomment if only Admin should access
		handlers.GetUsers,
	)

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

	// ==========================================================
	// MERCHANT ROUTES
	// ==========================================================
	// All APIs require JWT.
	// Update / Delete -> ADMIN Only
	// View -> Logged-in User

	

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

	// ==========================================================
	// TRANSACTION ROUTES (JWT Token Required)
	// ==========================================================
	// These APIs work only after successful login.

	router.POST("/transactions",
		middleware.JWTMiddleware(),
		handlers.CreateTransaction,
	)

	router.GET("/transactions",
    middleware.JWTMiddleware(),
    middleware.AdminMiddleware(),
    handlers.GetTransactions,
	)

	router.GET("/transactions/:id",
    middleware.JWTMiddleware(),
    middleware.AdminMiddleware(),
    handlers.GetTransactionByID,
)

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

	// ==========================================================
	// PAYBACK ROUTE (JWT Token Required)
	// ==========================================================
	// User must login before making a payback.

	router.POST("/payback",
		middleware.JWTMiddleware(),
		middleware.UserMiddleware(),
		handlers.Payback,
	)

	// ==========================================================
	// REPORT ROUTES
	// ==========================================================
	// These APIs require:
	// 1. Valid JWT Token
	// 2. ADMIN access only

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

	// ==========================================================
	// ADMIN ROUTES
	// ==========================================================
	// These APIs require:
	// 1. Valid JWT Token
	// 2. ADMIN access only

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