package config

import (
	"fmt"
	"os"
	"strings"
)

// LoadJWTSecret requires JWT_SECRET.
func LoadJWTSecret() (string, error) {
	return MustGetenv("JWT_SECRET")
}

// LoadInternalAPIToken requires INTERNAL_API_TOKEN (falls back to INTERNAL_TOKEN).
func LoadInternalAPIToken() (string, error) {
	if v := strings.TrimSpace(os.Getenv("INTERNAL_API_TOKEN")); v != "" {
		return v, nil
	}
	if v := strings.TrimSpace(os.Getenv("INTERNAL_TOKEN")); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("required environment variable INTERNAL_API_TOKEN is missing")
}

// ServicePort returns SERVICE_PORT or the provided default.
func ServicePort(fallback string) string {
	return Getenv("SERVICE_PORT", fallback)
}

// ServiceName returns SERVICE_NAME or the provided default.
func ServiceName(fallback string) string {
	return Getenv("SERVICE_NAME", fallback)
}

// Addr returns ":PORT" listen address.
func Addr(port string) string {
	port = strings.TrimPrefix(strings.TrimSpace(port), ":")
	return ":" + port
}
