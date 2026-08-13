//go:build ignore

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		fail(err)
	}

	pemPath := filepath.Join(root, "login_private.pem")
	if err := os.Remove(pemPath); err != nil && !os.IsNotExist(err) {
		fail(err)
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fail(err)
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err := os.WriteFile(pemPath, privateKeyPEM, 0o600); err != nil {
		fail(err)
	}

	escaped := strings.ReplaceAll(strings.TrimSpace(string(privateKeyPEM)), "\n", `\n`)
	if err := upsertEnv(filepath.Join(root, ".env"), "LOGIN_RSA_PRIVATE_KEY", escaped); err != nil {
		fail(err)
	}

	backendEnv := filepath.Join(root, "..", ".env")
	if err := upsertEnv(backendEnv, "LOGIN_RSA_PRIVATE_KEY", escaped); err != nil {
		fail(err)
	}

	fmt.Println("RSA login key configured in AuthService/.env, backend/.env, and login_private.pem.")
}

func upsertEnv(path, key, value string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
	prefix := key + "="
	updated := false

	for index, line := range lines {
		if strings.HasPrefix(line, prefix) {
			lines[index] = prefix + value
			updated = true
			break
		}
	}

	if !updated {
		if len(lines) > 0 && lines[len(lines)-1] != "" {
			lines = append(lines, "")
		}
		lines = append(lines, prefix+value)
	}

	output := strings.Join(lines, "\n")
	if !strings.HasSuffix(output, "\n") {
		output += "\n"
	}

	return os.WriteFile(path, []byte(output), 0o600)
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "setup failed:", err)
	os.Exit(1)
}
