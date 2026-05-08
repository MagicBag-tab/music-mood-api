package models

type Rating struct {
	ID        int    `json:"id,omitempty"`
	SongID    int    `json:"song_id"`
	Score     int    `json:"score"`
	Comment   string `json:"comment,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type RatingRequest struct {
	SongID  int    `json:"song_id"`
	Score   int    `json:"score"`
	Comment string `json:"comment,omitempty"`
}
