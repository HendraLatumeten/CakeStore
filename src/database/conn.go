package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Getconnection() (*sql.DB, error) {
	// db, err := sql.Open("mysql", "root:password@tcp(database:3306)/cakestoreDB")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakestoreDB")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)
	return db, nil
}
