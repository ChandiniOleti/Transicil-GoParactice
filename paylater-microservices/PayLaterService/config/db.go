package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() error {
	var err error

	dsn := "chandini:Chandini@123@tcp(localhost:3306)/paylaterdb?parseTime=true"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	fmt.Println("✅ Database Connected Successfully")

	return nil
}