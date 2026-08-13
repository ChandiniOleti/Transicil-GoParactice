package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"time"
)

const loginPasswordMaxAge = 5 * time.Minute

type encryptedPasswordPayload struct {
	Password string `json:"password"`
	Ts       int64  `json:"ts"`
}

var loginPrivateKey *rsa.PrivateKey

// InitLoginCrypto loads the RSA private key used to decrypt login passwords.
func InitLoginCrypto(privateKeyPEM string) error {
	normalizedPEM := strings.ReplaceAll(privateKeyPEM, `\n`, "\n")
	block, _ := pem.Decode([]byte(normalizedPEM))
	if block == nil {
		return errors.New("invalid LOGIN_RSA_PRIVATE_KEY PEM")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		pkcs1Key, pkcs1Err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if pkcs1Err != nil {
			return fmt.Errorf("parse login private key: %w", err)
		}
		loginPrivateKey = pkcs1Key
		return nil
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return errors.New("LOGIN_RSA_PRIVATE_KEY must be an RSA key")
	}

	loginPrivateKey = rsaKey
	return nil
}

// LoginPublicKeyPEM returns the PEM-encoded RSA public key for client encryption.
func LoginPublicKeyPEM() (string, error) {
	if loginPrivateKey == nil {
		return "", errors.New("login crypto is not initialized")
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&loginPrivateKey.PublicKey)
	if err != nil {
		return "", err
	}

	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})), nil
}

// DecryptLoginPassword decrypts an RSA-OAEP encrypted login password.
func DecryptLoginPassword(encryptedPassword string) (string, error) {
	if loginPrivateKey == nil {
		return "", errors.New("login crypto is not initialized")
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encryptedPassword)
	if err != nil {
		return "", errors.New("invalid encrypted password")
	}

	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		loginPrivateKey,
		ciphertext,
		nil,
	)
	if err != nil {
		return "", errors.New("invalid encrypted password")
	}

	var payload encryptedPasswordPayload
	if err := json.Unmarshal(plaintext, &payload); err != nil {
		return "", errors.New("invalid encrypted password")
	}

	if payload.Password == "" {
		return "", errors.New("invalid encrypted password")
	}

	if payload.Ts == 0 {
		return "", errors.New("invalid encrypted password")
	}

	issuedAt := time.Unix(payload.Ts, 0)
	now := time.Now()
	if issuedAt.After(now.Add(time.Minute)) || now.Sub(issuedAt) > loginPasswordMaxAge {
		return "", errors.New("encrypted password expired")
	}

	return payload.Password, nil
}
