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

func (r *SongRepository) Create(song models.SongRequest) (*models.Song, error) {
	query := `
		INSERT INTO songs (artist_id, album_id, title, mood, source, spotify_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, artist_id, album_id, title, mood, source, 
		          spotify_id, image_path, created_at, updated_at
	`

	var s models.Song
	err := r.db.QueryRow(query,
		song.ArtistID, song.AlbumID, song.Title,
		song.Mood, song.Source, song.SpotifyID,
	).Scan(
		&s.ID, &s.ArtistID, &s.AlbumID, &s.Title,
		&s.Mood, &s.Source, &s.SpotifyID,
		&s.ImagePath, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
