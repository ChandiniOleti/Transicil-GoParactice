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
	AuthServiceURL        string
	UserServiceURL        string
	MerchantServiceURL    string
	TransactionServiceURL string
	PaybackServiceURL     string
	ReportServiceURL      string
)

func Load() error {
	sharedconfig.LoadDotEnv()

	ServiceName = sharedconfig.ServiceName("APIGateway")
	Port = sharedconfig.ServicePort("8080")

	secret, err := sharedconfig.LoadJWTSecret()
	if err != nil {
		return err
	}
	JWTSecret = secret

	required := []string{
		"AUTH_SERVICE_URL",
		"USER_SERVICE_URL",
		"MERCHANT_SERVICE_URL",
		"TRANSACTION_SERVICE_URL",
		"PAYBACK_SERVICE_URL",
		"REPORT_SERVICE_URL",
	}
	if err := sharedconfig.Require(required...); err != nil {
		return err
	}

	AuthServiceURL = sharedconfig.Getenv("AUTH_SERVICE_URL", "")
	UserServiceURL = sharedconfig.Getenv("USER_SERVICE_URL", "")
	MerchantServiceURL = sharedconfig.Getenv("MERCHANT_SERVICE_URL", "")
	TransactionServiceURL = sharedconfig.Getenv("TRANSACTION_SERVICE_URL", "")
	PaybackServiceURL = sharedconfig.Getenv("PAYBACK_SERVICE_URL", "")
	ReportServiceURL = sharedconfig.Getenv("REPORT_SERVICE_URL", "")

	logger.L().Info("config loaded",
		"jwt_secret", sharedconfig.Mask(JWTSecret),
		"auth", AuthServiceURL,
		"user", UserServiceURL,
		"merchant", MerchantServiceURL,
		"transaction", TransactionServiceURL,
		"payback", PaybackServiceURL,
		"report", ReportServiceURL,
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
	return nil
}
