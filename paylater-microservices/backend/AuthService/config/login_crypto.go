package config

import (
	"fmt"
	"os"
	"strings"
)

// LoadLoginRSAPrivateKey reads the PEM-encoded RSA private key for login decryption.
func LoadLoginRSAPrivateKey() (string, error) {
	privateKeyPEM := strings.TrimSpace(os.Getenv("LOGIN_RSA_PRIVATE_KEY"))
	if privateKeyPEM == "" {
		return "", fmt.Errorf("LOGIN_RSA_PRIVATE_KEY is required")
	}

	return privateKeyPEM, nil
}
