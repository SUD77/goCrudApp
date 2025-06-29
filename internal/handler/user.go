package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/SUD77/goCrudApp/internal/model"
)

func CreateUser(db *sql.DB) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {

		var u model.User

		if err := json.NewDecoder(r.Body).Decode(&u); err!=nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := db.QueryRow(model.InserUser, u.Name, u.Email).Scan(&u.ID)

		if err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	}
}