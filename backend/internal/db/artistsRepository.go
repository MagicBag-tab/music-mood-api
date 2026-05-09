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

func (r *ArtistRepository) GetAll() ([]models.Artist, error) {
	rows, err := r.db.Query(`SELECT id, name, country, image_path, created_at FROM artists ORDER BY name ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artists []models.Artist
	for rows.Next() {
		var a models.Artist
		if err := rows.Scan(&a.ID, &a.Name, &a.Country, &a.ImagePath, &a.CreatedAt); err != nil {
			return nil, err
		}
		artists = append(artists, a)
	}
	return artists, nil
}

func (r *ArtistRepository) Create(req models.ArtistRequest) (*models.Artist, error) {
	var a models.Artist
	err := r.db.QueryRow(
		`INSERT INTO artists (name, country, image_path) 
         VALUES ($1, $2, $3) 
         RETURNING id, name, country, image_path, created_at`,
		req.Name, req.Country, req.ImagePath,
	).Scan(&a.ID, &a.Name, &a.Country, &a.ImagePath, &a.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *ArtistRepository) Update(id int, req models.ArtistRequest) (*models.Artist, error) {
	var a models.Artist
	err := r.db.QueryRow(
		`UPDATE artists SET name = $1, country = $2, image_path = $3 
         WHERE id = $4 
         RETURNING id, name, country, image_path, created_at`,
		req.Name, req.Country, req.ImagePath, id,
	).Scan(&a.ID, &a.Name, &a.Country, &a.ImagePath, &a.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *ArtistRepository) Delete(id int) error {
	res, err := r.db.Exec(`DELETE FROM artists WHERE id = $1`, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}
