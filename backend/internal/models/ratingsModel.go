package models

type Rating struct {
	ID        int    `json:"id,omitempty"`
	SongID    int    `json:"song_id"`
	Score     int    `json:"score"`
	Comment   string `json:"comment,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type RatingRequest struct {
	SongID  int    `json:"song_id" binding:"required"`
	Score   int    `json:"score" binding:"required,min=1,max=5"`
	Comment string `json:"comment,omitempty"`
}
