package models

import "time"

type Song struct {
	ID         int       `json:"id"`
	ArtistID   int       `json:"artist_id"`
	ArtistName string    `json:"artist_name,omitempty"`
	AlbumID    *int      `json:"album_id,omitempty"`
	Title      string    `json:"title"`
	Mood       string    `json:"mood"`
	Source     string    `json:"source"`
	SpotifyID  *string   `json:"spotify_id,omitempty"`
	ImagePath  *string   `json:"image_path,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type SongRequest struct {
	ArtistID  int    `json:"artist_id"`
	AlbumID   *int   `json:"album_id,omitempty"`
	Title     string `json:"title"`
	Mood      string `json:"mood"`
	Source    string `json:"source"`
	SpotifyID string `json:"spotify_id"`
}

type SongFilters struct {
	Page   int
	Limit  int
	Search string
	Sort   string
	Order  string
	Mood   string
}
