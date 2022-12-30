package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DB_DSN")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB Mysql Connected")
	}
}
