package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("paylater_secret_key")


func GenerateToken(userID int32, role string) (string, error) {

	claims := jwt.MapClaims{

		"user_id": userID,
		"role": role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}


	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)


	return token.SignedString(SecretKey)
}