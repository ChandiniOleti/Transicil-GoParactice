package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dsn := "chandini:Chandini@123@tcp(127.0.0.1:3306)/playersdb?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	}

	log.Println("Database connected successfully")

	return db
}
