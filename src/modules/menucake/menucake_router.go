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

	route.HandleFunc("/list", ctrl.ListAll).Methods("GET")
	route.HandleFunc("/detail/{id}", ctrl.DetailAll).Methods("GET")
	route.HandleFunc("/save", ctrl.Add).Methods("POST")

}
