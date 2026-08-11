package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

// Must be initialized via InitJWT from JWT_SECRET.
var SecretKey []byte

// InitJWT configures the signing key from environment configuration.
func InitJWT(secret string) {
	SecretKey = []byte(secret)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
}
