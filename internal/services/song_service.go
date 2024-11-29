package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"music_library/internal/models"
	"net/http"
	"strconv"
)

func GetSongs(db *sql.DB, group, song, page, limit string) ([]models.Song, error) {
	query := "SELECT id, group_name, song_name, release_date, text, link FROM songs"

	if group != "" {
		query += fmt.Sprintf(" AND group_name LIKE '%%%s%%'", group)
	}
	if song != "" {
		query += fmt.Sprintf(" AND song_name LIKE '%%%s%%'", song)
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	offset := (pageInt - 1) * limitInt
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limitInt, offset)

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
	if err := rows.Err(); err != nil {
		return nil, err
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

func GetSongText(db *sql.DB, songID, page, limit string) ([]string, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 1
	}

	offset := (pageInt - 1) * limitInt

	query := `SELECT text FROM song_lyrics WHERE song_id = $1 LIMIT $2 OFFSET $3`
	rows, err := db.Query(query, songID, limitInt, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var textParts []string
	for rows.Next() {
		var text string
		if err := rows.Scan(&text); err != nil {
			return nil, err
		}
		textParts = append(textParts, text)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return textParts, nil
}

func FetchSongDetailsFromAPI(groupName, songName string) (models.Song, error) {
	apiURL := fmt.Sprintf("http://external-api.com/info?group=%s&song=%s", groupName, songName)
	resp, err := http.Get(apiURL)
	if err != nil {
		return models.Song{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Song{}, fmt.Errorf("failed to fetch song details from API")
	}

	var songDetail models.Song
	err = json.NewDecoder(resp.Body).Decode(&songDetail)
	if err != nil {
		return models.Song{}, err
	}

	return songDetail, nil
}

func AddSongWithDetails(db *sql.DB, song models.Song, songDetail models.Song) (int, error) {
	query := `INSERT INTO songs (group_name, song_name, release_date, text, link) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int
	err := db.QueryRow(query, song.GroupName, song.SongName, songDetail.ReleaseDate, songDetail.Text, songDetail.Link).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteSong(db *sql.DB, songID string) error {
	_, err := db.Exec("DELETE FROM songs WHERE id = $1", songID)
	return err
}

func UpdateSong(db *sql.DB, song models.Song) error {
	_, err := db.Exec(
		"UPDATE songs SET group_name = $1, song_name = $2, release_date = $3, text = $4, link = $5 WHERE id = $6",
		song.GroupName, song.SongName, song.ReleaseDate, song.Text, song.Link, song.ID,
	)
	return err
}
