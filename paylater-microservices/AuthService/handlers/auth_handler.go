package handlers

import (
	"github.com/gin-gonic/gin"

	"authservice/services"
)

// LoginHandler authenticates a normal user and returns a JWT token.
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

// AdminLoginHandler authenticates an admin and returns a JWT token.
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

// MerchantLoginHandler authenticates a merchant and returns a JWT token.
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
