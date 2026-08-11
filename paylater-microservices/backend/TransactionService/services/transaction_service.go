package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"transactionservice/clients"
	"transactionservice/config"
	"transactionservice/generated"
)

// TransactionResponse is the clean API DTO for transaction history endpoints.
// Nullable SQLC fields are converted to plain JSON-friendly types.
type TransactionResponse struct {
	ID                   int32   `json:"id"`
	UserID               int32   `json:"user_id"`
	MerchantID           *int32  `json:"merchant_id,omitempty"`
	Amount               string  `json:"amount"`
	CommissionPercentage string  `json:"commission_percentage"`
	CommissionAmount     string  `json:"commission_amount"`
	TransactionType      string  `json:"transaction_type"`
	TransactionDate      *string `json:"transaction_date,omitempty"`
}

func toTransactionResponse(t generated.Transaction) TransactionResponse {
	response := TransactionResponse{
		ID:                   t.ID,
		UserID:               t.UserID,
		Amount:               t.Amount,
		CommissionPercentage: t.CommissionPercentage,
		CommissionAmount:     t.CommissionAmount,
		TransactionType:      string(t.TransactionType),
	}

	if t.MerchantID.Valid {
		merchantID := t.MerchantID.Int32
		response.MerchantID = &merchantID
	}

	if t.TransactionDate.Valid {
		formatted := t.TransactionDate.Time.UTC().Format(time.RFC3339)
		response.TransactionDate = &formatted
	}

	return response
}

func toTransactionResponses(transactions []generated.Transaction) []TransactionResponse {
	response := make([]TransactionResponse, 0, len(transactions))
	for _, transaction := range transactions {
		response = append(response, toTransactionResponse(transaction))
	}
	return response
}

// ProcessTransaction handles a purchase transaction via REST orchestration.
//
// Flow:
// 1. Get user details (User Service).
// 2. Get merchant commission (Merchant Service).
// 3. Check available credit.
// 4. Calculate merchant commission.
// 5. Save the purchase transaction.
// 6. Update the user's current due (User Service).
func ProcessTransaction(transaction generated.CreateTransactionParams, authHeader string) (map[string]interface{}, error) {
	ctx := context.Background()

	user, err := clients.GetUserByID(transaction.UserID, authHeader)
	if err != nil {
		return nil, err
	}

	merchant, err := clients.GetMerchantCommission(transaction.MerchantID.Int32)
	if err != nil {
		return nil, err
	}

	creditLimit, err := strconv.ParseFloat(user.CreditLimit, 64)
	if err != nil {
		return nil, err
	}

	currentDue, err := strconv.ParseFloat(user.CurrentDue, 64)
	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseFloat(transaction.Amount, 64)
	if err != nil {
		return nil, err
	}

	commissionPercent, err := strconv.ParseFloat(merchant.Commission, 64)
	if err != nil {
		return nil, err
	}

	availableCredit := creditLimit - currentDue

	if amount > availableCredit {
		return nil, errors.New("credit limit exceeded")
	}

	commissionAmount := (amount * commissionPercent) / 100

	transaction.CommissionPercentage = fmt.Sprintf("%.2f", commissionPercent)
	transaction.CommissionAmount = fmt.Sprintf("%.2f", commissionAmount)
	transaction.TransactionType = generated.TransactionsTransactionTypePURCHASE

	_, err = config.Queries.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	newDue := currentDue + amount

	err = clients.UpdateCurrentDue(user.ID, fmt.Sprintf("%.2f", newDue))
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"message":               "Transaction Successful",
		"user_id":               transaction.UserID,
		"merchant_id":           transaction.MerchantID.Int32,
		"transaction_type":      transaction.TransactionType,
		"transaction_amount":    transaction.Amount,
		"commission_percentage": fmt.Sprintf("%.2f", commissionPercent),
		"commission_amount":     fmt.Sprintf("%.2f", commissionAmount),
		"updated_current_due":   fmt.Sprintf("%.2f", newDue),
		"available_credit":      fmt.Sprintf("%.2f", creditLimit-newDue),
	}

	return response, nil
}

func GetTransactions() ([]TransactionResponse, error) {
	transactions, err := config.Queries.GetTransactions(context.Background())
	if err != nil {
		return nil, err
	}
	return toTransactionResponses(transactions), nil
}

func GetTransactionByID(id int32) (TransactionResponse, error) {
	transaction, err := config.Queries.GetTransactionByID(context.Background(), id)
	if err != nil {
		return TransactionResponse{}, err
	}
	return toTransactionResponse(transaction), nil
}

func GetTransactionsByUser(id int32) ([]TransactionResponse, error) {
	transactions, err := config.Queries.GetTransactionsByUser(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return toTransactionResponses(transactions), nil
}

func GetTransactionsByMerchant(id int32) ([]TransactionResponse, error) {
	transactions, err := config.Queries.GetTransactionsByMerchant(
		context.Background(),
		sql.NullInt32{
			Int32: id,
			Valid: true,
		},
	)
	if err != nil {
		return nil, err
	}
	return toTransactionResponses(transactions), nil
}

// CreatePaybackTransaction inserts a PAYBACK row (null merchant, zero commission).
func CreatePaybackTransaction(userID int32, amount string) (TransactionResponse, error) {
	ctx := context.Background()

	result, err := config.Queries.CreateTransaction(ctx, generated.CreateTransactionParams{
		UserID: userID,
		MerchantID: sql.NullInt32{
			Valid: false,
		},
		Amount:               amount,
		CommissionPercentage: "0.00",
		CommissionAmount:     "0.00",
		TransactionType:      generated.TransactionsTransactionTypePAYBACK,
	})
	if err != nil {
		return TransactionResponse{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return TransactionResponse{}, err
	}

	return GetTransactionByID(int32(id))
}
