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

		role := c.GetString("role")
		tokenMerchantID := c.GetInt("user_id")

		if role == "ADMIN" {
			c.Next()
			return
		}

		if role != "MERCHANT" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Merchant access only",
			})
			c.Abort()
			return
		}

		param := c.Param("merchant_id")

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
