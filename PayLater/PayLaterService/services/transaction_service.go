package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"paylaterservice/config"
	"paylaterservice/generated"
	"database/sql"
)

// Process Transaction (Purchase)
func ProcessTransaction(transaction generated.CreateTransactionParams) (map[string]interface{}, error) {

	ctx := context.Background()

	// Get User
	user, err := config.Queries.GetUserByID(ctx, transaction.UserID)
	if err != nil {
		return nil, err
	}

	// Get Merchant
	merchant, err := config.Queries.GetMerchantByID(ctx, transaction.MerchantID.Int32)
	if err != nil {
		return nil, err
	}

	// Convert Values
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

	// Available Credit
	availableCredit := creditLimit - currentDue

	if amount > availableCredit {
		return nil, errors.New("credit limit exceeded")
	}

	// Commission Calculation
	commissionAmount := (amount * commissionPercent) / 100

	// Set Transaction Values
	transaction.CommissionPercentage = fmt.Sprintf("%.2f", commissionPercent)
	transaction.CommissionAmount = fmt.Sprintf("%.2f", commissionAmount)
	transaction.TransactionType = generated.TransactionsTransactionTypePURCHASE

	// Insert Transaction
	_, err = config.Queries.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	// Update Current Due
	newDue := currentDue + amount

	err = config.Queries.UpdateCurrentDue(ctx, generated.UpdateCurrentDueParams{
		ID:         user.ID,
		CurrentDue: fmt.Sprintf("%.2f", newDue),
	})

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


// Get All Transactions
func GetTransactions() ([]generated.Transaction, error) {

	transactions, err := config.Queries.GetTransactions(context.Background())

	if err != nil {
		return nil, err
	}

	return transactions, nil
}


// Get Transaction By ID
func GetTransactionByID(id int32) (generated.Transaction, error) {

	transaction, err := config.Queries.GetTransactionByID(context.Background(), id)

	if err != nil {
		return generated.Transaction{}, err
	}

	return transaction, nil
}


// Get Transactions By User
func GetTransactionsByUser(id int32) ([]generated.Transaction, error) {

	transactions, err := config.Queries.GetTransactionsByUser(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}


// Get Transactions By Merchant
func GetTransactionsByMerchant(id int32) ([]generated.Transaction, error) {

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