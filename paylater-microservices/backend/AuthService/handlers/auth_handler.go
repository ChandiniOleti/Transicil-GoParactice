package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"authservice/services"
	"authservice/utils"
)

// EncryptedLoginRequest is the only accepted login request body.
// Email is sent in plaintext; only the password is RSA-OAEP encrypted.
type EncryptedLoginRequest struct {
	Email             string `json:"email" binding:"required"`
	EncryptedPassword string `json:"encrypted_password" binding:"required"`
}

func parseEncryptedLoginRequest(c *gin.Context) (services.LoginRequest, error) {
	var request EncryptedLoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		return services.LoginRequest{}, err
	}

	password, err := utils.DecryptLoginPassword(request.EncryptedPassword)
	if err != nil {
		return services.LoginRequest{}, err
	}

	return services.LoginRequest{
		Email:    request.Email,
		Password: password,
	}, nil
}

// LoginPublicKeyHandler exposes the RSA public key used for login password encryption.
func LoginPublicKeyHandler(c *gin.Context) {
	publicKeyPEM, err := utils.LoginPublicKeyPEM()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"public_key": publicKeyPEM,
	})
}

// LoginHandler authenticates a normal user and returns a JWT token.
func LoginHandler(c *gin.Context) {
	request, err := parseEncryptedLoginRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid login credentials",
		})
		return
	}

	token, err := services.Login(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Login Successful",
		"token":   token,
	})
}

// AdminLoginHandler authenticates an admin and returns a JWT token.
func AdminLoginHandler(c *gin.Context) {
	request, err := parseEncryptedLoginRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid login credentials",
		})
		return
	}

	token, err := services.AdminLogin(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin Login Successful",
		"token":   token,
	})
}

// MerchantLoginHandler authenticates a merchant and returns a JWT token.
func MerchantLoginHandler(c *gin.Context) {
	request, err := parseEncryptedLoginRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid login credentials",
		})
		return
	}

	token, err := services.MerchantLogin(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Merchant Login Successful",
		"token":   token,
	})
}
