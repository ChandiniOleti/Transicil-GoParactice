package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key used to sign and verify JWT tokens.
// The same key must be used while generating and validating the token.
var SecretKey = []byte("paylater_secret_key")

// GenerateToken creates a JWT token after successful login.
//
// Parameters:
// userID - Logged-in user's ID
// email  - User's email
// role   - User's role (ADMIN / USER)
//
// Example:
// User ID : 1
// Email   : admin@gmail.com
// Role    : ADMIN
//
// A token is generated with these details (claims)
// and is valid for 24 hours.
func GenerateToken(userID int32, email string, role string) (string, error) {

	// Store user information inside the JWT token.
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,

		// Token expiration time (24 hours from now)
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	// Create a new JWT token using the HS256 signing algorithm.
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	// Sign the token using the secret key and return it.
	return token.SignedString(SecretKey)
}

// ValidateToken verifies whether the received JWT token is valid.
//
// This function is called by the JWT middleware before
// allowing access to protected APIs.
//
// Example:
//
// Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
//
// If the token is valid, user details can be extracted
// from the claims.
func ValidateToken(tokenString string) (*jwt.Token, error) {

	// Parse the token and verify its signature
	// using the application's secret key.
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return SecretKey, nil
	})
}