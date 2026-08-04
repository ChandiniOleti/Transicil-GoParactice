package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"paybackservice/config"
)

var httpClient = &http.Client{Timeout: 5 * time.Second}

// UserDetails is the subset of user data needed for payback.
type UserDetails struct {
	ID          int32  `json:"id"`
	CreditLimit string `json:"credit_limit"`
	CurrentDue  string `json:"current_due"`
}

// GetUserByID fetches user details from User Service using the caller's JWT.
func GetUserByID(userID int32, authHeader string) (*UserDetails, error) {
	url := fmt.Sprintf("%s/users/%d", config.UserServiceURL, userID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", authHeader)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user service error (%d): %s", resp.StatusCode, string(body))
	}

	var user UserDetails
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateCurrentDue updates user due via User Service internal API.
func UpdateCurrentDue(userID int32, currentDue string) error {
	url := fmt.Sprintf("%s/internal/users/%d/due", config.UserServiceURL, userID)

	payload, err := json.Marshal(map[string]string{
		"current_due": currentDue,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Internal-Token", config.InternalToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("user due update failed (%d): %s", resp.StatusCode, string(body))
	}
	return nil
}
