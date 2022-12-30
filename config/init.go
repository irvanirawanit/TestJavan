package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

	// migrate from file javan.sql
	// execute the sql script from file
	err = loadSQLFile(DB, "./javan.sql")
	if err != nil {
		panic(err)
	}
	// sqlScript, err := ioutil.ReadFile("./javan.sql")
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = DB.Exec(string(sqlScript))
	// if err != nil {
	// 	panic(err)
	// }
	// DB.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), email VARCHAR(255))")
}

func loadSQLFile(db *sql.DB, sqlFile string) error {
	file, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		tx.Rollback()
	}()
	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if _, err := tx.Exec(q); err != nil {
			return err
		}
	}
	return tx.Commit()
}
