package models

import "time"

type Album struct {
	ID          int       `json:"id"`
	ArtistID    int       `json:"artist_id"`
	Title       string    `json:"title"`
	ReleaseDate string    `json:"release_date"`
	CoverPath   string    `json:"cover_path,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type AlbumRequest struct {
	ArtistID    int    `json:"artist_id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	CoverPath   string `json:"cover_path"`
}
