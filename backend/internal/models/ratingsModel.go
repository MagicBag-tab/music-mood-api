package models

import "time"

type Rating struct {
	ID        int       `json:"id"`
	SongID    int       `json:"song_id"`
	Score     int       `json:"score"`
	Comment   string    `json:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type RatingRequest struct {
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}
