package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// Config holds application configuration values
type Config struct {
    DatabaseURL string // Postgres connection string
}

// Load reads environment variables (via .env) and returns a Config
func Load() Config {
    // Load .env file if present
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  No .env file found; relying on existing environment variables")
    }

    // Read DATABASE_URL
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL must be set in environment or .env file")
    }

    return Config{
        DatabaseURL: dbURL,
    }
}
