package middleware

import (
	"github.com/gin-gonic/gin"
)

// AdminMiddleware allows only ADMIN role.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")

		if !exists {
			c.JSON(401, gin.H{
				"error": "Role not found",
			})
			c.Abort()
			return
		}

		if role != "ADMIN" {
			c.JSON(403, gin.H{
				"error": "Access Denied. Admin Only",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
