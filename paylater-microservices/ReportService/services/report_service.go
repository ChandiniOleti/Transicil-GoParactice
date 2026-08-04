package services

import (
	"fmt"
	"sort"
	"strconv"

	"reportservice/clients"
)

// MerchantFeeResponse represents the merchant commission report.
type MerchantFeeResponse struct {
	MerchantID        int32  `json:"merchant_id"`
	TotalFeeCollected string `json:"total_fee_collected"`
}

// TotalDuesResponse represents the total outstanding dues report.
type TotalDuesResponse struct {
	TotalDue string `json:"total_due"`
}

// GetMerchantFeeReport returns total commission collected by a merchant.
func GetMerchantFeeReport(merchantID int32) (interface{}, error) {
	transactions, err := clients.GetTransactionsByMerchant(merchantID)
	if err != nil {
		return nil, err
	}

	var total float64
	for _, tx := range transactions {
		if tx.TransactionType != "PURCHASE" {
			continue
		}
		fee, err := strconv.ParseFloat(tx.CommissionAmount, 64)
		if err != nil {
			return nil, err
		}
		total += fee
	}

	return MerchantFeeResponse{
		MerchantID:        merchantID,
		TotalFeeCollected: fmt.Sprintf("%.2f", total),
	}, nil
}

// GetUsersWithDueReport returns users with current_due > 0.
func GetUsersWithDueReport() (interface{}, error) {
	users, err := clients.GetUsers()
	if err != nil {
		return nil, err
	}

	result := make([]clients.UserReport, 0)
	for _, user := range users {
		due, err := strconv.ParseFloat(user.CurrentDue, 64)
		if err != nil {
			return nil, err
		}
		if due > 0 {
			result = append(result, user)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		di, _ := strconv.ParseFloat(result[i].CurrentDue, 64)
		dj, _ := strconv.ParseFloat(result[j].CurrentDue, 64)
		return di > dj
	})

	return result, nil
}

// GetUserDueReport returns one user's due details.
func GetUserDueReport(userID int32) (interface{}, error) {
	user, err := clients.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetCreditLimitUsersReport returns users who reached/exceeded credit limit.
func GetCreditLimitUsersReport() (interface{}, error) {
	users, err := clients.GetUsers()
	if err != nil {
		return nil, err
	}

	result := make([]clients.UserReport, 0)
	for _, user := range users {
		due, err := strconv.ParseFloat(user.CurrentDue, 64)
		if err != nil {
			return nil, err
		}
		limit, err := strconv.ParseFloat(user.CreditLimit, 64)
		if err != nil {
			return nil, err
		}
		if due >= limit {
			result = append(result, user)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		di, _ := strconv.ParseFloat(result[i].CurrentDue, 64)
		dj, _ := strconv.ParseFloat(result[j].CurrentDue, 64)
		return di > dj
	})

	return result, nil
}

// GetTotalDuesReport returns SUM of all users' current_due.
func GetTotalDuesReport() (interface{}, error) {
	users, err := clients.GetUsers()
	if err != nil {
		return nil, err
	}

	var total float64
	for _, user := range users {
		due, err := strconv.ParseFloat(user.CurrentDue, 64)
		if err != nil {
			return nil, err
		}
		total += due
	}

	return TotalDuesResponse{
		TotalDue: fmt.Sprintf("%.2f", total),
	}, nil
}
