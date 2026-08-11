package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"reportservice/config"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

// UserReport is the clean user DTO used by reports (no password).
type UserReport struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CreditLimit string `json:"credit_limit"`
	CurrentDue  string `json:"current_due"`
}

func GetUsers() ([]UserReport, error) {
	url := fmt.Sprintf("%s/internal/users", config.UserServiceURL)

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
		return nil, fmt.Errorf("user service error (%d): %s", resp.StatusCode, string(body))
	}

	var users []UserReport
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, err
	}
	if users == nil {
		users = []UserReport{}
	}
	return users, nil
}

func GetUserByID(userID int32) (*UserReport, error) {
	url := fmt.Sprintf("%s/internal/users/%d", config.UserServiceURL, userID)

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
		return nil, fmt.Errorf("user service error (%d): %s", resp.StatusCode, string(body))
	}

	var user UserReport
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}
	return &user, nil
}
