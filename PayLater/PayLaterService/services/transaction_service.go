package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"paylaterservice/config"
	"paylaterservice/generated"
)

// ProcessTransaction handles a purchase transaction.
//
// Flow:
// 1. Get user details.
// 2. Get merchant details.
// 3. Check available credit.
// 4. Calculate merchant commission.
// 5. Save the purchase transaction.
// 6. Update the user's current due.
//
// Example:
// Credit Limit : 2000
// Current Due  : 500
// Purchase     : 300
//
// New Due      : 800
// Available Credit : 1200
func ProcessTransaction(transaction generated.CreateTransactionParams) (map[string]interface{}, error) {

	ctx := context.Background()

	// Fetch user details using User ID.
	user, err := config.Queries.GetUserByID(ctx, transaction.UserID)
	if err != nil {
		return nil, err
	}

	// Fetch merchant details using Merchant ID.
	merchant, err := config.Queries.GetMerchantByID(ctx, transaction.MerchantID.Int32)
	if err != nil {
		return nil, err
	}

	// Convert string values to float for calculations.
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

	// Calculate available credit.
	availableCredit := creditLimit - currentDue

	// Do not allow purchase if the credit limit is exceeded.
	if amount > availableCredit {
		return nil, errors.New("credit limit exceeded")
	}

	// Calculate merchant commission.
	//
	// Formula:
	// Commission = Purchase Amount × Commission Percentage / 100
	commissionAmount := (amount * commissionPercent) / 100

	// Set commission values and transaction type.
	transaction.CommissionPercentage = fmt.Sprintf("%.2f", commissionPercent)
	transaction.CommissionAmount = fmt.Sprintf("%.2f", commissionAmount)
	transaction.TransactionType = generated.TransactionsTransactionTypePURCHASE

	// Save the purchase transaction into the transactions table.
	_, err = config.Queries.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	// Update user's current due after successful purchase.
	newDue := currentDue + amount

	err = config.Queries.UpdateCurrentDue(ctx, generated.UpdateCurrentDueParams{
		ID:         user.ID,
		CurrentDue: fmt.Sprintf("%.2f", newDue),
	})

	if err != nil {
		return nil, err
	}

	// Return transaction summary.
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

// GetTransactions returns all transactions.
//
// Example:
// GET /transactions
func GetTransactions() ([]generated.Transaction, error) {

	// Fetch all transaction records.
	transactions, err := config.Queries.GetTransactions(context.Background())

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetTransactionByID returns a single transaction using its ID.
//
// Example:
// GET /transactions/1
func GetTransactionByID(id int32) (generated.Transaction, error) {

	// Fetch transaction by ID.
	transaction, err := config.Queries.GetTransactionByID(context.Background(), id)

	if err != nil {
		return generated.Transaction{}, err
	}

	return transaction, nil
}

// GetTransactionsByUser returns all transactions made by a specific user.
//
// Example:
// GET /transactions/user/1
func GetTransactionsByUser(id int32) ([]generated.Transaction, error) {

	// Fetch transactions using User ID.
	transactions, err := config.Queries.GetTransactionsByUser(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetTransactionsByMerchant returns all purchase transactions
// performed through a specific merchant.
//
// Note:
// Payback transactions have MerchantID = NULL,
// so they will not be included in this report.
//
// Example:
// GET /transactions/merchant/2
func GetTransactionsByMerchant(id int32) ([]generated.Transaction, error) {

	// Convert Merchant ID into sql.NullInt32 because
	// MerchantID is nullable in the database.
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

	return transactions, nil
}