package db

import (
	"database/sql"
	"fmt"
	"music-mood-api/internal/models"
)

type SongRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) GetAll(f models.SongFilters) ([]models.Song, int, error) {
	where := "WHERE 1=1"
	args := []any{}
	i := 1

	if f.Search != "" {
		where += fmt.Sprintf(" AND (s.title ILIKE $%d OR a.name ILIKE $%d)", i, i+1)
		args = append(args, "%"+f.Search+"%", "%"+f.Search+"%")
		i += 2
	}

	if f.Mood != "" {
		where += fmt.Sprintf(" AND s.mood = $%d", i)
		args = append(args, f.Mood)
		i++
	}

	countQuery := fmt.Sprintf(`
		SELECT COUNT(*) FROM songs s
		JOIN artists a ON s.artist_id = a.id
		%s`, where)

	var total int
	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	validSorts := map[string]bool{"title": true, "created_at": true, "mood": true}
	sortCol := "s.created_at"
	if validSorts[f.Sort] {
		sortCol = "s." + f.Sort
	}
	order := "DESC"
	if f.Order == "asc" {
		order = "ASC"
	}

	if f.Limit <= 0 {
		f.Limit = 10
	}
	if f.Page <= 0 {
		f.Page = 1
	}
	offset := (f.Page - 1) * f.Limit

	args = append(args, f.Limit, offset)
	query := fmt.Sprintf(`
		SELECT s.id, s.artist_id, s.album_id, s.title, s.mood, s.source,
		       s.spotify_id, s.image_path, s.created_at, s.updated_at,
		       a.name as artist_name
		FROM songs s
		JOIN artists a ON s.artist_id = a.id
		%s
		ORDER BY %s %s
		LIMIT $%d OFFSET $%d`,
		where, sortCol, order, i, i+1)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var s models.Song
		err := rows.Scan(
			&s.ID, &s.ArtistID, &s.AlbumID, &s.Title,
			&s.Mood, &s.Source, &s.SpotifyID,
			&s.ImagePath, &s.CreatedAt, &s.UpdatedAt,
			&s.ArtistName,
		)
		if err != nil {
			return nil, 0, err
		}
		songs = append(songs, s)
	}

	if songs == nil {
		songs = []models.Song{}
	}
	return songs, total, nil
}

func (r *SongRepository) Create(song models.SongRequest) (*models.Song, error) {
	query := `
		INSERT INTO songs (artist_id, album_id, title, mood, source, spotify_id)
		VALUES ($1, $2, $3, $4, $5, NULLIF($6, ''))
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
		SELECT s.id, s.artist_id, s.album_id, s.title, s.mood, s.source,
		       s.spotify_id, s.image_path, s.created_at, s.updated_at,
		       a.name as artist_name
		FROM songs s
		JOIN artists a ON s.artist_id = a.id
		WHERE s.id = $1`, id,
	).Scan(
		&s.ID, &s.ArtistID, &s.AlbumID, &s.Title,
		&s.Mood, &s.Source, &s.SpotifyID,
		&s.ImagePath, &s.CreatedAt, &s.UpdatedAt,
		&s.ArtistName,
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
			spotify_id = NULLIF($6, ''),
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

func (r *SongRepository) UpdateImage(id int, imagePath string) error {
	_, err := r.db.Exec(
		`UPDATE songs SET image_path = $1, updated_at = NOW() WHERE id = $2`,
		imagePath, id,
	)
	return err
}
