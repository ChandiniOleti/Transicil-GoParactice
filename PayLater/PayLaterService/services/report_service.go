package services

import (
	"context"
	"database/sql"
	"fmt"

	"paylaterservice/config"
)

// MerchantFeeResponse represents the response for the
// merchant commission report.
type MerchantFeeResponse struct {
	MerchantID        int32  `json:"merchant_id"`
	TotalFeeCollected string `json:"total_fee_collected"`
}

// TotalDuesResponse represents the response for the
// total outstanding dues report.
type TotalDuesResponse struct {
	TotalDue string `json:"total_due"`
}

// GetMerchantFeeReport returns the total commission
// collected by a specific merchant.
//
// Example:
// GET /reports/merchant-fee/1
//
// Response:
// {
//     "merchant_id": 1,
//     "total_fee_collected": "150.00"
// }
func GetMerchantFeeReport(merchantID int32) (interface{}, error) {

	ctx := context.Background()

	// Fetch the total commission collected by the merchant.
	result, err := config.Queries.GetMerchantFeeCollected(
		ctx,
		sql.NullInt32{
			Int32: merchantID,
			Valid: true,
		},
	)

	if err != nil {
		return nil, err
	}

	// Convert the returned value into a string.
	totalFee := ""

	if result.TotalFeeCollected != nil {
		totalFee = string(result.TotalFeeCollected.([]byte))
	}

	// Prepare the API response.
	response := MerchantFeeResponse{
		MerchantID:        result.MerchantID.Int32,
		TotalFeeCollected: totalFee,
	}

	return response, nil
}

// GetUsersWithDueReport returns all users whose
// current due is greater than zero.
//
// Example:
// GET /reports/users-with-due
func GetUsersWithDueReport() (interface{}, error) {

	ctx := context.Background()

	// Fetch users who have pending dues.
	result, err := config.Queries.GetUsersWithDue(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetUserDueReport returns the current due
// of a specific user.
//
// Example:
// GET /reports/user-due/2
func GetUserDueReport(userID int32) (interface{}, error) {

	ctx := context.Background()

	// Fetch user details using User ID.
	result, err := config.Queries.GetUserDue(ctx, userID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCreditLimitUsersReport returns all users
// who have reached their credit limit.
//
// Example:
// GET /reports/credit-limit-users
func GetCreditLimitUsersReport() (interface{}, error) {

	ctx := context.Background()

	// Fetch users whose current due
	// is equal to or greater than their credit limit.
	result, err := config.Queries.GetUsersReachedCreditLimit(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTotalDuesReport returns the total outstanding
// due amount of all users.
//
// Example:
// GET /reports/total-dues
//
// Response:
// {
//     "total_due": "3500.00"
// }
func GetTotalDuesReport() (interface{}, error) {

	ctx := context.Background()

	// Fetch the total dues from the database.
	result, err := config.Queries.GetTotalUserDues(ctx)

	if err != nil {
		return nil, err
	}

	// Convert the returned value into a string.
	totalDue := ""

	switch value := result.(type) {

	case []byte:
		totalDue = string(value)

	case string:
		totalDue = value

	default:
		totalDue = fmt.Sprintf("%v", value)
	}

	// Prepare the API response.
	response := TotalDuesResponse{
		TotalDue: totalDue,
	}

	return response, nil
}