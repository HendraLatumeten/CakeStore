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

	route.HandleFunc("", ctrl.Test).Methods("GET")
	route.HandleFunc("/list", ctrl.ListAll).Methods("GET")
	route.HandleFunc("/detail/{id}", ctrl.DetailAll).Methods("GET")
	route.HandleFunc("/save", ctrl.AddData).Methods("POST")
	route.HandleFunc("/delete/{id}", ctrl.DeleteData).Methods("DELETE")
	route.HandleFunc("/update/{id}", ctrl.UpdateData).Methods("PUT")

}
