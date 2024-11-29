package handlers

import (
	"database/sql"
	"encoding/json"
	"music_library/internal/models"
	"music_library/internal/services"
	"music_library/internal/utils"
	"net/http"
)

func GetSongsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.InfoLogger.Println("Fetching songs...")

		songs, err := services.GetSongs(db, r.URL.Query())
		if err != nil {
			http.Error(w, "Failed to fetch songs", http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(songs)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func AddSongHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var song models.Song

		if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		id, err := services.AddSong(db, song)
		if err != nil {
			http.Error(w, "Failed to add song", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(map[string]int{"id": id})
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}
