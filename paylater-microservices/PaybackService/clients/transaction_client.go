package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"paybackservice/config"
)

// PaybackTransactionResponse is the clean DTO returned by Transaction Service.
type PaybackTransactionResponse struct {
	ID                   int32   `json:"id"`
	UserID               int32   `json:"user_id"`
	MerchantID           *int32  `json:"merchant_id,omitempty"`
	Amount               string  `json:"amount"`
	CommissionPercentage string  `json:"commission_percentage"`
	CommissionAmount     string  `json:"commission_amount"`
	TransactionType      string  `json:"transaction_type"`
	TransactionDate      *string `json:"transaction_date,omitempty"`
}

// CreatePaybackTransaction records a PAYBACK row via Transaction Service.
func CreatePaybackTransaction(userID int32, amount string) (*PaybackTransactionResponse, error) {
	url := fmt.Sprintf("%s/internal/transactions/payback", config.TransactionServiceURL)

	payload, err := json.Marshal(map[string]interface{}{
		"user_id": userID,
		"amount":  amount,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Internal-Token", config.InternalToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("transaction service error (%d): %s", resp.StatusCode, string(body))
	}

	var transaction PaybackTransactionResponse
	if err := json.Unmarshal(body, &transaction); err != nil {
		return nil, err
	}
	return &transaction, nil
}
