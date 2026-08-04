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
// ==========================================================

// SecretKey is used to digitally sign and verify JWTs.
// Must be initialized via InitJWT from JWT_SECRET.
var SecretKey []byte

// InitJWT configures the signing key from environment configuration.
func InitJWT(secret string) {
	SecretKey = []byte(secret)
}

// GenerateToken creates a JWT token after successful login.
func GenerateToken(userID int32, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(SecretKey)
}

// ValidateToken validates the received JWT token.
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
}
