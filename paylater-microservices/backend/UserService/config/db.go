package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	sharedconfig "paylater.dev/shared/config"
	"paylater.dev/shared/logger"
)

var DB *sql.DB

var (
	ServiceName   string
	Port          string
	JWTSecret     string
	InternalToken string
)

func Load() error {
	sharedconfig.LoadDotEnv()

	ServiceName = sharedconfig.ServiceName("UserService")
	Port = sharedconfig.ServicePort("8082")

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
	return nil
}

func ConnectDB() error {
	dbCfg, err := sharedconfig.LoadDBConfig()
	if err != nil {
		return err
	}

	DB, err = sql.Open("mysql", dbCfg.DSN())
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}

	logger.L().Info("database connected",
		"host", dbCfg.Host,
		"port", dbCfg.Port,
		"db", dbCfg.Name,
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
