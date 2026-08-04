package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	
)

// UserMiddleware allows:
//
// 1. ADMIN -> Can access any user's data.
// 2. USER  -> Can access only their own data.
// 3. MERCHANT -> Cannot access user APIs.
func UserMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Get role and user ID from JWT
		role := c.GetString("role")
		tokenUserID := c.GetInt("user_id")

		// Admin can access everything
		if role == "ADMIN" {
			c.Next()
			return
		}

		// Only USER role is allowed after this point
		if role != "USER" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "User access only",
			})
			c.Abort()
			return
		}

		// Get User ID from URL
		param := c.Param("user_id")

		// For routes like /users/:id
		if param == "" {
			param = c.Param("id")
		}

		paramID, err := strconv.Atoi(param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid User ID",
			})
			c.Abort()
			return
		}

		// User can access only their own account
		if tokenUserID != paramID {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "You can access only your own account",
			})
			c.Abort()
			return
		}

		c.Next()
	}
	
}