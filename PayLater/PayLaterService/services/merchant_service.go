package services

import (
	"context"
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"paylaterservice/config"
	"paylaterservice/generated"
)

// merchant_service.go contains the business logic for all
// merchant-related operations.
//
// Authorization:
// - Merchant : Register
// - ADMIN    : Update, Update Commission, Delete
// - USER     : View merchant details only

// CreateMerchant adds a new merchant to the database.
//
// Authorization:
// Public API (No JWT Required)
//
// Business Rule:
// Merchant can register only if the commission
// is between 3% and 10%.
func CreateMerchant(merchant generated.CreateMerchantParams) error {

	// Validate commission
	commission, err := strconv.ParseFloat(merchant.Commission, 64)
	if err != nil {
		return errors.New("invalid commission")
	}

	if commission < 3 || commission > 10 {
		return errors.New("commission must be between 3 and 10")
	}

	// Encrypt password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(merchant.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	merchant.Password = string(hashedPassword)

	// Insert merchant into database
	_, err = config.Queries.CreateMerchant(context.Background(), merchant)

	if err != nil {
		return err
	}

	return nil
}

// GetMerchants returns all merchants.
//
// Authorization:
// ADMIN and USER
func GetMerchants() ([]generated.Merchant, error) {

	merchants, err := config.Queries.GetMerchants(context.Background())

	if err != nil {
		return nil, err
	}

	return merchants, nil
}

// GetMerchantByID returns one merchant.
//
// Authorization:
// Merchant (Own Profile)
// ADMIN
func GetMerchantByID(id int32) (generated.Merchant, error) {

	merchant, err := config.Queries.GetMerchantByID(
		context.Background(),
		id,
	)

	if err != nil {
		return generated.Merchant{}, err
	}

	return merchant, nil
}

// UpdateMerchant updates merchant details.
//
// Authorization:
// ADMIN Only
func UpdateMerchant(merchant generated.UpdateMerchantParams) error {

	// Encrypt updated password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(merchant.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	merchant.Password = string(hashedPassword)

	err = config.Queries.UpdateMerchant(
		context.Background(),
		merchant,
	)

	if err != nil {
		return err
	}

	return nil
}

// UpdateCommission updates only commission.
//
// Authorization:
// ADMIN Only
//
// NOTE:
// Admin can set ANY commission.
// No validation here.
func UpdateCommission(merchant generated.UpdateCommissionParams) error {

	return config.Queries.UpdateCommission(
		context.Background(),
		merchant,
	)
}

// DeleteMerchant deletes a merchant.
//
// Authorization:
// ADMIN Only
func DeleteMerchant(id int32) error {

	return config.Queries.DeleteMerchant(
		context.Background(),
		id,
	)
}