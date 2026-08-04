package middleware

import (
	"github.com/gin-gonic/gin"
)

// ==========================================================
// ADMIN MIDDLEWARE
//
// This middleware allows only users with the ADMIN role
// to access protected admin APIs.
//
// Responsibilities:
// 1. Read the user's role from Gin Context.
// 2. Check whether the role exists.
// 3. Verify the role is ADMIN.
// 4. Allow the request if authorized.
// 5. Stop the request if unauthorized.
// ==========================================================
func AdminMiddleware() gin.HandlerFunc {

	// Return a middleware function that executes before
	// the actual handler.
	return func(c *gin.Context) {

		// Read the user's role from Gin Context.
		// The role was stored by JWT Authentication Middleware.
		role, exists := c.Get("role")

		// Check whether the role exists.
		if !exists {

			// Return 401 Unauthorized if role is missing.
			c.JSON(401, gin.H{
				"error": "Role not found",
			})

			// Stop executing the remaining handlers.
			c.Abort()
			return
		}

		// Verify that the logged-in user is an ADMIN.
		if role != "ADMIN" {

			// Return 403 Forbidden if the user is not an admin.
			c.JSON(403, gin.H{
				"error": "Access Denied. Admin Only",
			})

			// Stop the request.
			c.Abort()
			return
		}

		// User is authorized.
		// Continue to the next middleware or handler.
		c.Next()
	}
}