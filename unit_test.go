package main

import (
	"testing"

	"github.com/hendralatumeten/cakestore/src/database"
	"github.com/hendralatumeten/cakestore/src/database/models"
	"github.com/hendralatumeten/cakestore/src/libs"
)

func TestGetConnection(t *testing.T) {
	db, err := database.Getconnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Errorf("Error pinging database: %v", err)
	}
}
func TestGetResponse(t *testing.T) {
	var data models.MenuCake

	result := libs.Respone(data, 200, false)
	if result.Data != data {
		panic("data must be false")
	} else if result.Code != 200 {
		panic("code must be false")
	} else if result.IsError != false {
		panic("IsError must be false")
	}
}
