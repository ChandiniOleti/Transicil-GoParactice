package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		tokenUserID := c.GetInt("user_id")

		if role == "ADMIN" {
			c.Next()
			return
		}

		if role != "USER" {
			c.JSON(http.StatusForbidden, gin.H{"error": "User access only"})
			c.Abort()
			return
		}

		param := c.Param("user_id")
		if param == "" {
			param = c.Param("id")
		}

		paramID, err := strconv.Atoi(param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
			c.Abort()
			return
		}

		if tokenUserID != paramID {
			c.JSON(http.StatusForbidden, gin.H{"error": "You can access only your own account"})
			c.Abort()
			return
		}

		c.Next()
	}
}
