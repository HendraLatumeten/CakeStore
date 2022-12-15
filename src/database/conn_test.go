package database_test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConn(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakestoreDB")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
