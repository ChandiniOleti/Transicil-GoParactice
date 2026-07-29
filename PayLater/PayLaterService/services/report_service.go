package services

import (
    "context"
    "database/sql"

    "paylaterservice/config"
	"fmt"
)

type MerchantFeeResponse struct {
	MerchantID       int32  `json:"merchant_id"`
	TotalFeeCollected string `json:"total_fee_collected"`
}

type TotalDuesResponse struct {
	TotalDue string `json:"total_due"`
}

func GetMerchantFeeReport(merchantID int32) (interface{}, error) {

	ctx := context.Background()

	result, err := config.Queries.GetMerchantFeeCollected(ctx, sql.NullInt32{
		Int32: merchantID,
		Valid: true,
	})

	if err != nil {
		return nil, err
	}


	totalFee := ""

	if result.TotalFeeCollected != nil {
		totalFee = string(result.TotalFeeCollected.([]byte))
	}


	response := MerchantFeeResponse{
		MerchantID: result.MerchantID.Int32,
		TotalFeeCollected: totalFee,
	}


	return response, nil
}


func GetUsersWithDueReport() (interface{}, error) {

	ctx := context.Background()

	result, err := config.Queries.GetUsersWithDue(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}


func GetUserDueReport(userID int32) (interface{}, error) {

	ctx := context.Background()

	result, err := config.Queries.GetUserDue(ctx, userID)

	if err != nil {
		return nil, err
	}

	return result, nil
}


func GetCreditLimitUsersReport() (interface{}, error) {

	ctx := context.Background()

	result, err := config.Queries.GetUsersReachedCreditLimit(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTotalDuesReport() (interface{}, error) {

	ctx := context.Background()

	result, err := config.Queries.GetTotalUserDues(ctx)

	if err != nil {
		return nil, err
	}


	totalDue := ""

	switch value := result.(type) {

	case []byte:
		totalDue = string(value)

	case string:
		totalDue = value

	default:
		totalDue = fmt.Sprintf("%v", value)
	}


	response := TotalDuesResponse{
		TotalDue: totalDue,
	}


	return response, nil
}