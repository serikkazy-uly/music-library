package main

import (
	"database/sql"
	"log"
	"net/http"

	"music_library/internal/db"
	"music_library/internal/handlers"
	"music_library/internal/utils"
)

func main() {
	utils.InitLogger()

	utils.InfoLogger.Println("Starting application...")

	utils.LoadEnv()

	dbConn := db.Connect()
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {

		}
	}(dbConn)

	db.RunMigrations(dbConn)

	router := handlers.NewRouter(dbConn)

	port := utils.GetEnv("PORT", "8080")
	utils.InfoLogger.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
