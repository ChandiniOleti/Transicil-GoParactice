package middleware

import (
	"github.com/gin-gonic/gin"

	"transactionservice/config"
)

// InternalMiddleware allows only trusted internal service calls.
func InternalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Internal-Token")
		if token == "" || token != config.InternalToken {
			c.JSON(401, gin.H{
				"error": "Unauthorized internal request",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
