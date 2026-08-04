package services

import (
	"errors"
	"fmt"
	"strconv"

	"paybackservice/clients"
)

// PaybackRequest represents the request body for the Payback API.
type PaybackRequest struct {
	UserID int32  `json:"user_id"`
	Amount string `json:"amount"`
}

// ProcessPayback handles the payback process via REST orchestration.
//
// Flow:
// 1. Get user details (User Service).
// 2. Validate the payback amount.
// 3. Reduce the user's current due (User Service).
// 4. Save the payback transaction (Transaction Service).
// 5. Return the updated due and available credit.
func ProcessPayback(request PaybackRequest, authHeader string) (map[string]interface{}, error) {
	user, err := clients.GetUserByID(request.UserID, authHeader)
	if err != nil {
		return nil, err
	}

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

	if paybackAmount <= 0 {
		return nil, errors.New("invalid payback amount")
	}

	if paybackAmount > currentDue {
		return nil, errors.New("payback amount exceeds current due")
	}

	newDue := currentDue - paybackAmount

	err = clients.UpdateCurrentDue(user.ID, fmt.Sprintf("%.2f", newDue))
	if err != nil {
		return nil, err
	}

	_, err = clients.CreatePaybackTransaction(request.UserID, request.Amount)
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"message":             "Payback Successful",
		"user_id":             request.UserID,
		"amount_paid":         request.Amount,
		"updated_current_due": fmt.Sprintf("%.2f", newDue),
		"available_credit":    fmt.Sprintf("%.2f", creditLimit-newDue),
	}

	return response, nil
}
