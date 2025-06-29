package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_"github.com/lib/pq"
	"github.com/joho/godotenv"
)

func init() {

	//1. Load .env into environment
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, continuing with existing env vars")
	}
}

func main() {

	//2. Read db url fromm env
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Database_URL is not set. Please check your .env")
	}

	//3. Open the db connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	defer db.Close()

	// verify connectivity
	if err := db.Ping(); err != nil {
		log.Fatal("cannot png db:", err)
	}

	r := mux.NewRouter()

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
