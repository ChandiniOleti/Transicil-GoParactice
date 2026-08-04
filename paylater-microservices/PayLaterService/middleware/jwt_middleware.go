package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"paylaterservice/utils"
)

// ==========================================================
// JWT AUTHENTICATION MIDDLEWARE
//
// This middleware is responsible for:
// 1. Reading the JWT token from the Authorization header.
// 2. Validating the JWT token.
// 3. Extracting user information from the token.
// 4. Storing user information in Gin Context.
// 5. Allowing only authenticated users to access protected APIs.
// ==========================================================
func JWTMiddleware() gin.HandlerFunc {

	// Return a middleware function that executes before
	// the actual API handler.
	return func(c *gin.Context) {

		// Read the Authorization header.
		// Example:
		// Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
		authHeader := c.GetHeader("Authorization")

		// Check whether the Authorization header exists.
		if authHeader == "" {

			// Return 401 if token is missing.
			c.JSON(401, gin.H{
				"error": "Authorization header missing",
			})

			// Stop further request processing.
			c.Abort()
			return
		}

		// Remove the "Bearer " prefix from the header
		// to get the actual JWT token.
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the JWT token.
		token, err := utils.ValidateToken(tokenString)

		// Check whether the token is invalid or expired.
		if err != nil || !token.Valid {

			// Return 401 Unauthorized.
			c.JSON(401, gin.H{
				"error": "Invalid Token",
			})

			// Stop request processing.
			c.Abort()
			return
		}

		// Extract claims (payload) from the JWT token.
		claims := token.Claims.(jwt.MapClaims)

		// Store the authenticated user's ID in Gin Context.
		c.Set("user_id", int(claims["user_id"].(float64)))

		// Store the user's role in Gin Context.
		c.Set("role", claims["role"].(string))

		// Continue to the next middleware or handler.
		c.Next()
	}
}