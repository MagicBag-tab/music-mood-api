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

func (r *SongRepository) GetByID(id int) (*models.Song, error) {
	var s models.Song
	err := r.db.QueryRow(`
		SELECT id, artist_id, album_id, title, mood, source,
		       spotify_id, image_path, created_at, updated_at
		FROM songs WHERE id = $1`, id,
	).Scan(
		&s.ID, &s.ArtistID, &s.AlbumID, &s.Title,
		&s.Mood, &s.Source, &s.SpotifyID,
		&s.ImagePath, &s.CreatedAt, &s.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *SongRepository) Update(id int, req models.SongRequest) (*models.Song, error) {
	var s models.Song
	err := r.db.QueryRow(`
		UPDATE songs SET
			artist_id  = $1,
			album_id   = $2,
			title      = $3,
			mood       = $4,
			source     = $5,
			spotify_id = $6,
			updated_at = NOW()
		WHERE id = $7
		RETURNING id, artist_id, album_id, title, mood, source,
		          spotify_id, image_path, created_at, updated_at`,
		req.ArtistID, req.AlbumID, req.Title,
		req.Mood, req.Source, req.SpotifyID, id,
	).Scan(
		&s.ID, &s.ArtistID, &s.AlbumID, &s.Title,
		&s.Mood, &s.Source, &s.SpotifyID,
		&s.ImagePath, &s.CreatedAt, &s.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *SongRepository) Delete(id int) error {
	result, err := r.db.Exec(`DELETE FROM songs WHERE id = $1`, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
