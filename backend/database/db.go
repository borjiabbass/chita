package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	var err error
	dsn := "root:password@tcp(127.0.0.1:3306)/vpn_db" // یوزر/پسورد رو عوض کن
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.Ping()
}
