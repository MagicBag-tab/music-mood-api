package db

import (
	"database/sql"
	"music-mood-api/internal/models"
)

type ArtistRepository struct {
	db *sql.DB
}

func NewArtistRepository(db *sql.DB) *ArtistRepository {
	return &ArtistRepository{db: db}
}

func (r *ArtistRepository) GetByID(id int) (*models.Artist, error) {
	var a models.Artist
	err := r.db.QueryRow(
		`SELECT id, name, country, image_path, created_at 
         FROM artists WHERE id = $1`, id,
	).Scan(&a.ID, &a.Name, &a.Country, &a.ImagePath, &a.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &a, nil
}
