package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// LoadDotEnv loads a .env file from the current working directory if present.
func LoadDotEnv() {
	_ = godotenv.Load()
}

// Getenv returns an environment variable or a default value.
func Getenv(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

// MustGetenv returns an environment variable or panics/fails via error.
func MustGetenv(key string) (string, error) {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return "", fmt.Errorf("required environment variable %s is missing", key)
	}
	return v, nil
}

// Require fails if any of the keys are missing or empty.
func Require(keys ...string) error {
	var missing []string
	for _, key := range keys {
		if strings.TrimSpace(os.Getenv(key)) == "" {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}
	return nil
}

// GetenvInt returns an int env var or default.
func GetenvInt(key string, fallback int) int {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}

// DBConfig holds MySQL connection settings.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// LoadDBConfig reads DB_* variables (required).
func LoadDBConfig() (DBConfig, error) {
	if err := Require("DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"); err != nil {
		return DBConfig{}, err
	}
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}, nil
}

// DSN builds a MySQL DSN with parseTime=true.
func (c DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.User, c.Password, c.Host, c.Port, c.Name)
}

// Mask masks a secret for logging.
func Mask(value string) string {
	if value == "" {
		return ""
	}
	if len(value) <= 4 {
		return "****"
	}
	return value[:2] + "****" + value[len(value)-2:]
}
