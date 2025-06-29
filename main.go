package main

import (
	"log"
	"net/http"

	"github.com/SUD77/goCrudApp/internal/config"
	"github.com/SUD77/goCrudApp/internal/db"
	"github.com/SUD77/goCrudApp/internal/router"
)

func main() {
	// 1) Load and validate environment (loads .env)
	cfg := config.Load()

	// 2) Connect to the database
	conn, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("DB setup failed: %v", err)
	}
	defer conn.Close()

	// 3) Wire up HTTP routes
	r := router.Setup(conn)

	// 4) Start the server
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
