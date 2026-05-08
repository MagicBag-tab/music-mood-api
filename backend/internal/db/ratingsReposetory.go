package db

import (
	"database/sql"
	"music-mood-api/internal/models"
)

type RatingRepository struct {
	db *sql.DB
}

func NewRatingRepository(db *sql.DB) *RatingRepository {
	return &RatingRepository{db: db}
}

func (r *RatingRepository) Create(songID int, req models.RatingRequest) (*models.Rating, error) {
	var rating models.Rating
	err := r.db.QueryRow(`
		INSERT INTO ratings (song_id, score, comment)
		VALUES ($1, $2, $3)
		RETURNING id, song_id, score, comment, created_at`,
		songID, req.Score, req.Comment,
	).Scan(&rating.ID, &rating.SongID, &rating.Score, &rating.Comment, &rating.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

func (r *RatingRepository) GetBySongID(songID int) ([]models.Rating, error) {
	rows, err := r.db.Query(`
		SELECT id, song_id, score, comment, created_at
		FROM ratings WHERE song_id = $1
		ORDER BY created_at DESC`, songID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []models.Rating
	for rows.Next() {
		var rt models.Rating
		if err := rows.Scan(&rt.ID, &rt.SongID, &rt.Score, &rt.Comment, &rt.CreatedAt); err != nil {
			return nil, err
		}
		ratings = append(ratings, rt)
	}
	return ratings, nil
}

func (r *RatingRepository) GetAverageBySongID(songID int) (float64, error) {
	var avg float64
	err := r.db.QueryRow(`
		SELECT COALESCE(AVG(score), 0)
		FROM ratings WHERE song_id = $1`, songID,
	).Scan(&avg)
	return avg, err
}
