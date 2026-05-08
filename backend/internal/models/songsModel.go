package models

import "time"

type Song struct {
	ID        int       `json:"id"`
	ArtistID  int       `json:"artist_id"`
	AlbumID   *int      `json:"album_id,omitempty"`
	Title     string    `json:"title"`
	Mood      string    `json:"mood"`
	Source    string    `json:"source"`
	Genres    []string  `json:"genres,omitempty"`
	SpotifyID *string   `json:"spotify_id,omitempty"`
	ImagePath *string   `json:"image_path,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SongRequest struct {
	Title     string  `json:"title"`
	ArtistID  int     `json:"artist_id"`
	AlbumID   *int    `json:"album_id,omitempty"`
	Mood      string  `json:"mood"`
	Source    string  `json:"source"`
	GenreIDs  []int   `json:"genre_ids,omitempty"`
	SpotifyID *string `json:"spotify_id,omitempty"`
}
