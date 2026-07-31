package services

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"paylaterservice/config"
	"paylaterservice/utils"
)

// LoginRequest represents the request body for all Login APIs.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ======================================================
// USER LOGIN
// ======================================================

// Login authenticates a normal USER.
func Login(request LoginRequest) (string, error) {

	ctx := context.Background()

	// Find user by email.
	user, err := config.Queries.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	// Compare entered password with hashed password.
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	)

	if err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT token.
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

// ======================================================
// ADMIN LOGIN
// ======================================================

// AdminLogin authenticates an ADMIN.
func AdminLogin(request LoginRequest) (string, error) {

	ctx := context.Background()

	// Find admin by email.
	admin, err := config.Queries.GetAdminByEmail(ctx, request.Email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	// Compare entered password with hashed password.
	err = bcrypt.CompareHashAndPassword(
		[]byte(admin.Password),
		[]byte(request.Password),
	)

	if err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT token.
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

// ======================================================
// MERCHANT LOGIN
// ======================================================

// MerchantLogin authenticates a MERCHANT.
func MerchantLogin(request LoginRequest) (string, error) {

	ctx := context.Background()

	// Find merchant by email.
	merchant, err := config.Queries.GetMerchantByEmail(ctx, request.Email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	// Compare entered password with hashed password.
	err = bcrypt.CompareHashAndPassword(
		[]byte(merchant.Password),
		[]byte(request.Password),
	)

	if err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT token.
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