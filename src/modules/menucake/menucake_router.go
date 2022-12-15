package menucake

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// akan memangil semua method
// inisialisasi endpoint

func New(rt *mux.Router, db *sql.DB) {
	route := rt.PathPrefix("/menucake").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")

}
