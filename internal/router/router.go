package router

import (
	"database/sql"

	"github.com/SUD77/goCrudApp/internal/handler"
	"github.com/gorilla/mux"
)

func Setup(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", handler.CreateUser(db)).Methods("POST")
	r.HandleFunc("/users", handler.ListUsers(db)).Methods("GET")
	// r.HandleFunc("/users/{id}", handler.GetUser(db)).Methods("GET")
	// r.HandleFunc("/users/{id}", handler.UpdateUser(db)).Methods("PUT")
	// r.HandleFunc("/users/{id}", handler.DeleteUser(db)).Methods("DELETE")

	return r

}
