package router

import (
	"database/sql"

	"github.com/SUD77/goCrudApp/internal/handler"
	"github.com/gorilla/mux"
)

func Setup(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", handler.CreateUser(db)).Methods("POST")

	return r

}
