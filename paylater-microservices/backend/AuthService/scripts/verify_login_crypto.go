//go:build ignore

package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const gatewayBaseURL = "http://localhost:8080"

func main() {
	statuses := make([]string, 0, 4)

	publicKeyPEM, err := fetchPublicKey()
	if err != nil {
		fail(err)
	}
	statuses = append(statuses, "GET /login/public-key: OK")

	plaintextStatus, err := postJSON(gatewayBaseURL+"/login", map[string]string{
		"email":    "probe@example.com",
		"password": "probe-password",
	})
	if err != nil {
		fail(err)
	}
	if plaintextStatus != http.StatusBadRequest {
		fail(fmt.Errorf("plaintext login expected 400, got %d", plaintextStatus))
	}
	statuses = append(statuses, "plaintext login rejected: OK")

	encryptedBody, err := buildEncryptedBody(publicKeyPEM, "missing-user@example.com", "probe-password")
	if err != nil {
		fail(err)
	}
	if strings.Contains(encryptedBody, "probe-password") {
		fail(fmt.Errorf("encrypted body must not contain plaintext password"))
	}
	if !strings.Contains(encryptedBody, "encrypted_password") {
		fail(fmt.Errorf("encrypted body missing encrypted_password field"))
	}
	if !strings.Contains(encryptedBody, "missing-user@example.com") {
		fail(fmt.Errorf("encrypted body must include plaintext email"))
	}

	encryptedStatus, err := postRaw(gatewayBaseURL+"/login", encryptedBody)
	if err != nil {
		fail(err)
	}
	if encryptedStatus != http.StatusUnauthorized {
		fail(fmt.Errorf("encrypted invalid login expected 401, got %d", encryptedStatus))
	}
	statuses = append(statuses, "encrypted login accepted/decrypted: OK")

	email := strings.TrimSpace(os.Getenv("LOGIN_TEST_EMAIL"))
	password := os.Getenv("LOGIN_TEST_PASSWORD")
	if email != "" && password != "" {
		validBody, err := buildEncryptedBody(publicKeyPEM, email, password)
		if err != nil {
			fail(err)
		}
		validStatus, err := postRaw(gatewayBaseURL+"/login", validBody)
		if err != nil {
			fail(err)
		}
		if validStatus != http.StatusOK {
			fail(fmt.Errorf("valid encrypted login expected 200, got %d", validStatus))
		}
		statuses = append(statuses, "valid encrypted login: OK")
	} else {
		statuses = append(statuses, "valid encrypted login: skipped (set LOGIN_TEST_EMAIL and LOGIN_TEST_PASSWORD to enable)")
	}

	fmt.Println(strings.Join(statuses, "\n"))
}

func fetchPublicKey() (string, error) {
	response, err := http.Get(gatewayBaseURL + "/login/public-key")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("public key request returned %d", response.StatusCode)
	}

	var payload struct {
		PublicKey string `json:"public_key"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return "", err
	}
	if payload.PublicKey == "" {
		return "", fmt.Errorf("public key response missing public_key")
	}

	return payload.PublicKey, nil
}

func buildEncryptedBody(publicKeyPEM, email, password string) (string, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return "", fmt.Errorf("invalid public key PEM")
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("public key is not RSA")
	}

	payload, err := json.Marshal(map[string]any{
		"password": password,
		"ts":       time.Now().Unix(),
	})
	if err != nil {
		return "", err
	}

	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		payload,
		nil,
	)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(map[string]string{
		"email":              email,
		"encrypted_password": base64.StdEncoding.EncodeToString(ciphertext),
	})
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func postJSON(url string, payload map[string]string) (int, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}
	return postRaw(url, string(body))
}

func postRaw(url, body string) (int, error) {
	response, err := http.Post(url, "application/json", bytes.NewBufferString(body))
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	_, _ = io.Copy(io.Discard, response.Body)
	return response.StatusCode, nil
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "verification failed:", err)
	os.Exit(1)
}
