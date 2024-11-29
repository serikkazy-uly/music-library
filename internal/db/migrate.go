package db

import (
	"database/sql"
	"log"
	"os"
)

func RunMigrations(db *sql.DB) {
	migrationFile := "internal/db/migrations/001_create_songs_table.sql"

	migration, err := os.ReadFile(migrationFile)
	if err != nil {
		log.Fatalf("Failed to read migration file: %v", err)
	}

	_, err = db.Exec(string(migration))
	if err != nil {
		log.Fatalf("Failed to execute migration: %v", err)
	}

	log.Println("Migrations executed successfully")
}
