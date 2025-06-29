package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // register Postgres driver
)

// Connect opens a DB connection using the given Postgres DSN, verifies it, and returns *sql.DB.
func Connect(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}
	// Ping to verify the connection is alive.
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}
	return db, nil
}
