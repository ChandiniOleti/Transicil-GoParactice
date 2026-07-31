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

// PaybackRequest represents the request body
// for the Payback API.
//
// Example:
// {
//     "user_id": 1,
//     "amount": "500.00"
// }
type PaybackRequest struct {
	UserID int32  `json:"user_id"`
	Amount string `json:"amount"`
}

// ProcessPayback handles the payback process.
//
// Flow:
// 1. Get user details.
// 2. Validate the payback amount.
// 3. Reduce the user's current due.
// 4. Save the payback transaction.
// 5. Return the updated due and available credit.
//
// Example:
//
// Current Due      : 1300
// Payback Amount   : 300
//
// New Current Due  : 1000
// Available Credit : 1000
func ProcessPayback(request PaybackRequest) (map[string]interface{}, error) {

	ctx := context.Background()

	// Fetch user details using User ID.
	user, err := config.Queries.GetUserByID(ctx, request.UserID)

	if err != nil {
		return nil, err
	}

	// Convert current due from string to float
	// for calculation.
	currentDue, err := strconv.ParseFloat(user.CurrentDue, 64)

	if err != nil {
		return nil, err
	}

	// Convert credit limit from string to float.
	creditLimit, err := strconv.ParseFloat(user.CreditLimit, 64)

	if err != nil {
		return nil, err
	}

	// Convert payback amount from string to float.
	paybackAmount, err := strconv.ParseFloat(request.Amount, 64)

	if err != nil {
		return nil, err
	}

	// Validate that the payback amount is greater than zero.
	if paybackAmount <= 0 {
		return nil, errors.New("invalid payback amount")
	}

	// Prevent paying more than the current due.
	if paybackAmount > currentDue {
		return nil, errors.New("payback amount exceeds current due")
	}

	// Calculate the new due after payback.
	newDue := currentDue - paybackAmount

	// Update the user's current due in the users table.
	err = config.Queries.UpdateCurrentDue(ctx, generated.UpdateCurrentDueParams{
		ID:         user.ID,
		CurrentDue: fmt.Sprintf("%.2f", newDue),
	})

	if err != nil {
		return nil, err
	}

	// Save the payback transaction.
	//
	// MerchantID is NULL because a payback
	// is not associated with any merchant.
	//
	// Commission values are 0 because
	// merchants do not receive commission
	// on payback transactions.
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

	// Prepare the response sent back to the client.
	response := map[string]interface{}{

		"message": "Payback Successful",

		"user_id": request.UserID,

		"amount_paid": request.Amount,

		"updated_current_due": fmt.Sprintf("%.2f", newDue),

		"available_credit": fmt.Sprintf("%.2f", creditLimit-newDue),
	}

	return response, nil
}