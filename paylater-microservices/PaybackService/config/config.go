package config

import (
	"fmt"

	sharedconfig "paylater.dev/shared/config"
	"paylater.dev/shared/logger"
)

var (
	ServiceName           string
	Port                  string
	JWTSecret             string
	InternalToken         string
	UserServiceURL        string
	TransactionServiceURL string
)

func Load() error {
	sharedconfig.LoadDotEnv()

	ServiceName = sharedconfig.ServiceName("PaybackService")
	Port = sharedconfig.ServicePort("8085")

	secret, err := sharedconfig.LoadJWTSecret()
	if err != nil {
		return err
	}
	JWTSecret = secret

	token, err := sharedconfig.LoadInternalAPIToken()
	if err != nil {
		return err
	}
	InternalToken = token

	UserServiceURL = sharedconfig.Getenv("USER_SERVICE_URL", "")
	TransactionServiceURL = sharedconfig.Getenv("TRANSACTION_SERVICE_URL", "")
	if err := sharedconfig.Require("USER_SERVICE_URL", "TRANSACTION_SERVICE_URL"); err != nil {
		return err
	}

	logger.L().Info("config loaded",
		"user_service", UserServiceURL,
		"transaction_service", TransactionServiceURL,
		"jwt_secret", sharedconfig.Mask(JWTSecret),
		"internal_token", sharedconfig.Mask(InternalToken),
	)
	return nil
}

func ListenAddr() string {
	return sharedconfig.Addr(Port)
}

func Validate() error {
	if JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}
	if InternalToken == "" {
		return fmt.Errorf("INTERNAL_API_TOKEN is required")
	}
	return nil
}
