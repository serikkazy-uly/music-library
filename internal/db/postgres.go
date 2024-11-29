package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"music_library/internal/utils"
	"os"
)

func Connect() *sql.DB {
	dsn := os.Getenv("DB_DSN")
	utils.InfoLogger.Printf("Connecting to database with DSN: %s", dsn)
	log.Printf("Connecting to database with DSN: %s", os.Getenv("DB_DSN"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		utils.ErrorLogger.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.Ping(); err != nil {
		utils.ErrorLogger.Fatalf("Database ping failed: %v", err)
	}

	utils.InfoLogger.Println("Successfully connected to the database")
	return db
}
