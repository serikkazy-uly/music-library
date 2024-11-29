package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"music_library/internal/models"
	"music_library/internal/services"
	"music_library/internal/utils"
	"net/http"
)

// @Summary Get all songs
// @Description Retrieve a list of all songs in the library
// @Tags songs
// @Accept json
// @Produce json
// @Success 200 {array} models.Song
// @Router /songs [get]
func GetSongsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.InfoLogger.Println("Fetching songs...")

		group := r.URL.Query().Get("group")
		song := r.URL.Query().Get("song")
		page := r.URL.Query().Get("page")
		limit := r.URL.Query().Get("limit")

		songs, err := services.GetSongs(db, group, song, page, limit)
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

func GetSongTextHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		songID := r.URL.Query().Get("song_id")
		page := r.URL.Query().Get("page")
		limit := r.URL.Query().Get("limit")

		text, err := services.GetSongText(db, songID, page, limit)
		if err != nil {
			http.Error(w, "Failed to fetch song text", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(text)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func AddSongWithAPIHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var song models.Song

		if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		songDetail, err := services.FetchSongDetailsFromAPI(song.GroupName, song.SongName)
		if err != nil {
			http.Error(w, "Failed to fetch song details from external API", http.StatusInternalServerError)
			return
		}

		id, err := services.AddSongWithDetails(db, song, songDetail)
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

func DeleteSongHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		songID := r.URL.Query().Get("song_id")

		err := services.DeleteSong(db, songID)
		if err != nil {
			http.Error(w, "Failed to delete song", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func UpdateSongHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		songID := vars["id"]

		var song models.Song

		if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		song.ID = id

		err := services.UpdateSong(db, song)
		if err != nil {
			http.Error(w, "Failed to update song", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
