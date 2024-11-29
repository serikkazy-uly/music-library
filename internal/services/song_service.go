package services

import (
	"database/sql"
	"music_library/internal/models"
	"net/url"
)

func GetSongs(db *sql.DB, params url.Values) ([]models.Song, error) {
	query := "SELECT id, group_name, song_name, release_date, text, link FROM songs"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.GroupName, &song.SongName, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func AddSong(db *sql.DB, song models.Song) (int, error) {
	query := `INSERT INTO songs (group_name, song_name, release_date, text, link)
              VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var id int
	err := db.QueryRow(query, song.GroupName, song.SongName, song.ReleaseDate, song.Text, song.Link).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
