package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"reportservice/config"
)

// TransactionReport is the clean transaction DTO used by reports.
type TransactionReport struct {
	ID                   int32   `json:"id"`
	UserID               int32   `json:"user_id"`
	MerchantID           *int32  `json:"merchant_id,omitempty"`
	Amount               string  `json:"amount"`
	CommissionPercentage string  `json:"commission_percentage"`
	CommissionAmount     string  `json:"commission_amount"`
	TransactionType      string  `json:"transaction_type"`
	TransactionDate      *string `json:"transaction_date,omitempty"`
}

func GetTransactionsByMerchant(merchantID int32) ([]TransactionReport, error) {
	url := fmt.Sprintf("%s/internal/transactions/merchant/%d", config.TransactionServiceURL, merchantID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Internal-Token", config.InternalToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("transaction service error (%d): %s", resp.StatusCode, string(body))
	}

	var transactions []TransactionReport
	if err := json.Unmarshal(body, &transactions); err != nil {
		return nil, err
	}
	if transactions == nil {
		transactions = []TransactionReport{}
	}
	return transactions, nil
}
