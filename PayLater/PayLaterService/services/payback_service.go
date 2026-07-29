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

type PaybackRequest struct {
	UserID int32  `json:"user_id"`
	Amount string `json:"amount"`
}

func ProcessPayback(request PaybackRequest) (map[string]interface{}, error) {

	ctx := context.Background()

	// Get User
	user, err := config.Queries.GetUserByID(ctx, request.UserID)

	if err != nil {
		return nil, err
	}


	// Convert Values
	currentDue, err := strconv.ParseFloat(user.CurrentDue, 64)

	if err != nil {
		return nil, err
	}


	creditLimit, err := strconv.ParseFloat(user.CreditLimit, 64)

	if err != nil {
		return nil, err
	}


	paybackAmount, err := strconv.ParseFloat(request.Amount, 64)

	if err != nil {
		return nil, err
	}


	// Validation
	if paybackAmount <= 0 {
		return nil, errors.New("invalid payback amount")
	}


	if paybackAmount > currentDue {
		return nil, errors.New("payback amount exceeds current due")
	}


	// Calculate New Due
	newDue := currentDue - paybackAmount



	// Update User Current Due
	err = config.Queries.UpdateCurrentDue(ctx, generated.UpdateCurrentDueParams{

		ID: user.ID,

		CurrentDue: fmt.Sprintf("%.2f", newDue),
	})


	if err != nil {
		return nil, err
	}



	// Insert PAYBACK Transaction
	_, err = config.Queries.CreateTransaction(ctx, generated.CreateTransactionParams{

		UserID: request.UserID,

		MerchantID: sql.NullInt32{
			Valid: false,
		},

		Amount: request.Amount,

		CommissionPercentage: "0.00",

		CommissionAmount: "0.00",

		TransactionType: generated.TransactionsTransactionTypePAYBACK,
	})


	if err != nil {
		return nil, err
	}



	response := map[string]interface{}{

		"message": "Payback Successful",

		"user_id": request.UserID,

		"amount_paid": request.Amount,

		"updated_current_due": fmt.Sprintf("%.2f", newDue),

		"available_credit": fmt.Sprintf("%.2f", creditLimit-newDue),
	}


	return response, nil
}