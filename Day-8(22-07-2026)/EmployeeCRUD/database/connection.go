package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {

	cfg := mysql.Config{
		User:                 "gouser",
		Passwd:               "Go@123",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "golangpractice",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}