package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MerchantMiddleware allows:
// 1. ADMIN -> Can access any merchant.
// 2. MERCHANT -> Can access only their own account.
func MerchantMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Get role and merchant ID stored by JWTMiddleware
		role := c.GetString("role")
		tokenMerchantID := c.GetInt("user_id")

		// Admin can access everything
		if role == "ADMIN" {
			c.Next()
			return
		}

		// Only merchants are allowed
		if role != "MERCHANT" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Merchant access only",
			})
			c.Abort()
			return
		}

		// Get Merchant ID from URL
		param := c.Param("merchant_id")

		// If merchant_id is not present,
		// use id (for /merchants/:id)
		if param == "" {
			param = c.Param("id")
		}

		paramID, err := strconv.Atoi(param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Merchant ID",
			})
			c.Abort()
			return
		}

		// Merchant can access only their own account
		if tokenMerchantID != paramID {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "You can access only your own account",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}