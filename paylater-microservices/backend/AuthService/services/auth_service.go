package services

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"authservice/config"
	"authservice/utils"
)

// LoginRequest represents the request body for all Login APIs.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login authenticates a normal USER.
func Login(request LoginRequest) (string, error) {
	ctx := context.Background()

	user, err := config.Queries.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	)
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(
		user.ID,
		user.Email,
		"USER",
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

// AdminLogin authenticates an ADMIN.
func AdminLogin(request LoginRequest) (string, error) {
	ctx := context.Background()

	admin, err := config.Queries.GetAdminByEmail(ctx, request.Email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(admin.Password),
		[]byte(request.Password),
	)
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(
		admin.ID,
		admin.Email,
		"ADMIN",
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

// MerchantLogin authenticates a MERCHANT.
func MerchantLogin(request LoginRequest) (string, error) {
	ctx := context.Background()

	merchant, err := config.Queries.GetMerchantByEmail(ctx, request.Email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(merchant.Password),
		[]byte(request.Password),
	)
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(
		merchant.ID,
		merchant.Email,
		"MERCHANT",
	)
	if err != nil {
		return "", err
	}

	return token, nil
}
