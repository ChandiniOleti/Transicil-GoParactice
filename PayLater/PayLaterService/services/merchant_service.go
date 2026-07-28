package services

import (
	"context"

	"paylaterservice/config"
	"paylaterservice/generated"
)

func CreateMerchant(merchant generated.CreateMerchantParams) error {

	_, err := config.Queries.CreateMerchant(context.Background(), merchant)

	if err != nil {
		return err
	}

	return nil
}

func GetMerchants() ([]generated.Merchant, error) {

	merchants, err := config.Queries.GetMerchants(context.Background())

	if err != nil {
		return nil, err
	}

	return merchants, nil
}

func GetMerchantByID(id int32) (generated.Merchant, error) {

	merchant, err := config.Queries.GetMerchantByID(context.Background(), id)

	if err != nil {
		return generated.Merchant{}, err
	}

	return merchant, nil
}

func UpdateMerchant(merchant generated.UpdateMerchantParams) error {

	err := config.Queries.UpdateMerchant(context.Background(), merchant)

	if err != nil {
		return err
	}

	return nil
}

func UpdateCommission(merchant generated.UpdateCommissionParams) error {

	err := config.Queries.UpdateCommission(context.Background(), merchant)

	if err != nil {
		return err
	}

	return nil
}

func DeleteMerchant(id int32) error {

	err := config.Queries.DeleteMerchant(context.Background(), id)

	if err != nil {
		return err
	}

	return nil
}