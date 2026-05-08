package db

import (
	"database/sql"
	"music-mood-api/internal/models"
)

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (r *AlbumRepository) GetAll() ([]models.Album, error) {
	rows, err := r.db.Query(`SELECT id, artist_id, title, release_date, cover_path, created_at FROM albums`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []models.Album
	for rows.Next() {
		var a models.Album
		if err := rows.Scan(&a.ID, &a.ArtistID, &a.Title, &a.ReleaseDate, &a.CoverPath, &a.CreatedAt); err != nil {
			return nil, err
		}
		albums = append(albums, a)
	}
	return albums, nil
}

func (r *AlbumRepository) GetByID(id int) (*models.Album, error) {
	var a models.Album
	err := r.db.QueryRow(
		`SELECT id, artist_id, title, release_date, cover_path, created_at 
         FROM albums WHERE id = $1`, id,
	).Scan(&a.ID, &a.ArtistID, &a.Title, &a.ReleaseDate, &a.CoverPath, &a.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AlbumRepository) Create(req models.AlbumRequest) (*models.Album, error) {
	var a models.Album
	err := r.db.QueryRow(
		`INSERT INTO albums (artist_id, title, release_date, cover_path) 
         VALUES ($1, $2, $3, $4) 
         RETURNING id, artist_id, title, release_date, cover_path, created_at`,
		req.ArtistID, req.Title, req.ReleaseDate, req.CoverPath,
	).Scan(&a.ID, &a.ArtistID, &a.Title, &a.ReleaseDate, &a.CoverPath, &a.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AlbumRepository) Update(id int, req models.AlbumRequest) (*models.Album, error) {
	var a models.Album
	err := r.db.QueryRow(
		`UPDATE albums SET artist_id = $1, title = $2, release_date = $3, cover_path = $4 
         WHERE id = $5 
         RETURNING id, artist_id, title, release_date, cover_path, created_at`,
		req.ArtistID, req.Title, req.ReleaseDate, req.CoverPath, id,
	).Scan(&a.ID, &a.ArtistID, &a.Title, &a.ReleaseDate, &a.CoverPath, &a.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AlbumRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM albums WHERE id = $1`, id)
	return err
}
