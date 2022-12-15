package routers

import (
	"github.com/gorilla/mux"
	"github.com/hendralatumeten/cakestore/src/database"
	"github.com/hendralatumeten/cakestore/src/modules/menucake"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	db := database.Getconnection()
	menucake.New(mainRoute, db)

	return mainRoute, nil
}
