package routers

import (
	"errors"

	"github.com/gorilla/mux"
	"github.com/hendralatumeten/cakestore/src/database"
	"github.com/hendralatumeten/cakestore/src/modules/menucake"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	db, err := database.Getconnection()
	if err != nil {
		return nil, errors.New("gagal init database")
	}

	menucake.New(mainRoute, db)

	return mainRoute, nil
}
