package handlers

import (
	"github.com/gin-gonic/gin"

	"paylaterservice/services"
)

// ======================================================
// USER LOGIN
// POST /login
// ======================================================

func LoginHandler(c *gin.Context) {

	var request services.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := services.Login(request)

	if err != nil {

		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User Login Successful",
		"token":   token,
	})
}

// ======================================================
// ADMIN LOGIN
// POST /admin/login
// ======================================================

func AdminLoginHandler(c *gin.Context) {

	var request services.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := services.AdminLogin(request)

	if err != nil {

		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Admin Login Successful",
		"token":   token,
	})
}

// ======================================================
// MERCHANT LOGIN
// POST /merchant/login
// ======================================================

func MerchantLoginHandler(c *gin.Context) {

	var request services.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := services.MerchantLogin(request)

	if err != nil {

		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Merchant Login Successful",
		"token":   token,
	})
}