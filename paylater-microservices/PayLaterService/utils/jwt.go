package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ==========================================================
// JWT UTILITY
//
// This file is responsible for:
// 1. Generating JWT tokens after successful login.
// 2. Validating JWT tokens for protected APIs.
// 3. Storing authenticated user information inside the token.
//
// JWT (JSON Web Token) is used for authentication so that
// users do not need to log in for every request.
// ==========================================================

// ==========================================================
// Secret Key
//
// SecretKey is used to:
// 1. Digitally sign the JWT while generating it.
// 2. Verify the JWT signature during validation.
//
// Both GenerateToken() and ValidateToken() must use the
// same secret key.
//
// Note:
// In production, never hardcode the secret key.
// Store it securely using environment variables or
// a secret management service.
// ==========================================================
var SecretKey = []byte("paylater_secret_key")

// ==========================================================
// GenerateToken
//
// Creates a JWT token after successful login.
//
// Parameters:
// userID - Logged-in user's ID
// email  - User's email
// role   - User's role (ADMIN / USER / MERCHANT)
//
// Returns:
// JWT Token
// Error (if any)
//
// Example:
// User ID : 1
// Email   : admin@gmail.com
// Role    : ADMIN
//
// The generated token remains valid for 24 hours.
// ==========================================================
func GenerateToken(userID int32, email string, role string) (string, error) {

	// Create JWT claims (payload).
	// These values can later be extracted from the token
	// by the JWT Middleware.
	claims := jwt.MapClaims{

		// Logged-in user's ID.
		"user_id": userID,

		// Logged-in user's email.
		"email": email,

		// Logged-in user's role.
		// Example: ADMIN, USER, MERCHANT
		"role": role,

		// Token expiration time.
		// After 24 hours the token becomes invalid.
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	// Create a new JWT token using the HS256
	// (HMAC SHA-256) signing algorithm and attach
	// the claims to the token.
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	// Sign the JWT using the application's secret key
	// and return the generated token.
	return token.SignedString(SecretKey)
}

// ==========================================================
// ValidateToken
//
// Validates the received JWT token.
//
// This function is called by JWT Middleware before
// allowing access to protected APIs.
//
// Parameters:
// tokenString - JWT received from the Authorization header.
//
// Example:
//
// Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
//
// Returns:
// Parsed JWT Token
// Error (if validation fails)
// ==========================================================
func ValidateToken(tokenString string) (*jwt.Token, error) {

	// Parse the JWT token.
	//
	// During parsing, JWT automatically checks:
	// • Signature
	// • Token format
	// • Expiration time (exp claim)
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Provide the secret key required to verify
		// the JWT signature.
		//
		// If the signature matches and the token has
		// not expired, the token is considered valid.
		return SecretKey, nil
	})
}