package services

import (
	"context"
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"merchantservice/config"
	"merchantservice/generated"
)

// MerchantResponse is returned to clients without the password hash.
type MerchantResponse struct {
	ID           int32  `json:"id"`
	MerchantName string `json:"merchant_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Commission   string `json:"commission"`
}

func toMerchantResponse(m generated.Merchant) MerchantResponse {
	return MerchantResponse{
		ID:           m.ID,
		MerchantName: m.MerchantName,
		Email:        m.Email,
		Phone:        m.Phone,
		Commission:   m.Commission,
	}
}

// CreateMerchant adds a new merchant to the database.
//
// Business Rule:
// Merchant can register only if the commission
// is between 3% and 10%.
func CreateMerchant(merchant generated.CreateMerchantParams) error {
	commission, err := strconv.ParseFloat(merchant.Commission, 64)
	if err != nil {
		return errors.New("invalid commission")
	}

	if commission < 3 || commission > 10 {
		return errors.New("commission must be between 3 and 10")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(merchant.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	merchant.Password = string(hashedPassword)

	_, err = config.Queries.CreateMerchant(context.Background(), merchant)
	return err
}

// GetMerchants returns all merchants without password hashes.
func GetMerchants() ([]MerchantResponse, error) {
	merchants, err := config.Queries.GetMerchants(context.Background())
	if err != nil {
		return nil, err
	}

	response := make([]MerchantResponse, 0, len(merchants))
	for _, merchant := range merchants {
		response = append(response, toMerchantResponse(merchant))
	}

	return response, nil
}

// GetMerchantByID returns one merchant without the password hash.
func GetMerchantByID(id int32) (MerchantResponse, error) {
	merchant, err := config.Queries.GetMerchantByID(context.Background(), id)
	if err != nil {
		return MerchantResponse{}, err
	}

	return toMerchantResponse(merchant), nil
}

// UpdateMerchant updates merchant details.
func UpdateMerchant(merchant generated.UpdateMerchantParams) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(merchant.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	merchant.Password = string(hashedPassword)

	return config.Queries.UpdateMerchant(context.Background(), merchant)
}

// UpdateCommission updates only commission.
// Admin can set ANY commission (no validation).
func UpdateCommission(merchant generated.UpdateCommissionParams) error {
	return config.Queries.UpdateCommission(context.Background(), merchant)
}

// DeleteMerchant deletes a merchant.
func DeleteMerchant(id int32) error {
	return config.Queries.DeleteMerchant(context.Background(), id)
}

// CommissionResponse is returned by the internal commission API.
type CommissionResponse struct {
	MerchantID int32  `json:"merchant_id"`
	Commission string `json:"commission"`
}

// GetMerchantCommission returns only merchant_id and commission.
func GetMerchantCommission(id int32) (CommissionResponse, error) {
	merchant, err := config.Queries.GetMerchantByID(context.Background(), id)
	if err != nil {
		return CommissionResponse{}, err
	}

	return CommissionResponse{
		MerchantID: merchant.ID,
		Commission: merchant.Commission,
	}, nil
}
