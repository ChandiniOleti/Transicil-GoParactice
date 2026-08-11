package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"transactionservice/config"
)

// MerchantCommission is returned by Merchant Service internal API.
type MerchantCommission struct {
	MerchantID int32  `json:"merchant_id"`
	Commission string `json:"commission"`
}

// GetMerchantCommission fetches commission from Merchant Service.
func GetMerchantCommission(merchantID int32) (*MerchantCommission, error) {
	url := fmt.Sprintf("%s/internal/merchants/%d/commission", config.MerchantServiceURL, merchantID)

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
		return nil, fmt.Errorf("merchant service error (%d): %s", resp.StatusCode, string(body))
	}

	var commission MerchantCommission
	if err := json.Unmarshal(body, &commission); err != nil {
		return nil, err
	}
	return &commission, nil
}
