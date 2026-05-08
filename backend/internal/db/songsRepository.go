package db

import (
	"database/sql"
	"music-mood-api/internal/models"
)

type SongRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) GetAll() ([]models.Song, error) {
	query := `
		SELECT id, artist_id, album_id, title, mood, source, 
		       spotify_id, image_path, created_at, updated_at
		FROM songs
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var s models.Song
		err := rows.Scan(
			&s.ID, &s.ArtistID, &s.AlbumID, &s.Title,
			&s.Mood, &s.Source, &s.SpotifyID,
			&s.ImagePath, &s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		songs = append(songs, s)
	}

	return songs, nil
}
