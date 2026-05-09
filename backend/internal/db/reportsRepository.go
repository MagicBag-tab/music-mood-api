package db

import (
	"database/sql"
	"music-mood-api/internal/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetMoodDistribution() ([]models.MoodDistribution, error) {
	query := `
		SELECT mood, COUNT(*) 
		FROM songs 
		GROUP BY mood
		ORDER BY COUNT(*) DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var distribution []models.MoodDistribution
	for rows.Next() {
		var md models.MoodDistribution
		if err := rows.Scan(&md.Mood, &md.Count); err != nil {
			return nil, err
		}
		distribution = append(distribution, md)
	}
	return distribution, nil
}

func (r *ReportRepository) GetTopRatedSongs(limit int) ([]models.TopRatedSong, error) {
	query := `
		SELECT s.id, s.title, a.name, COALESCE(AVG(rt.score), 0) as avg_rating, COUNT(rt.id) as rating_count
		FROM songs s
		JOIN artists a ON s.artist_id = a.id
		LEFT JOIN ratings rt ON s.id = rt.song_id
		GROUP BY s.id, a.name
		ORDER BY avg_rating DESC, rating_count DESC
		LIMIT $1
	`
	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topSongs []models.TopRatedSong
	for rows.Next() {
		var s models.TopRatedSong
		if err := rows.Scan(&s.SongID, &s.Title, &s.ArtistName, &s.AverageRating, &s.RatingsCount); err != nil {
			return nil, err
		}
		topSongs = append(topSongs, s)
	}
	return topSongs, nil
}
