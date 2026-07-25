package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {

	var err error

	DB, err = sql.Open(
		"mysql",
		"chandini:Chandini@123@tcp(localhost:3306)/employee_db",
	)

	if err != nil {
		panic(err)
	}

	err = DB.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("✅ MySQL Connected Successfully")
}