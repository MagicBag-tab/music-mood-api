package models

type Song struct {
	ID        int      `json:"id"`
	ArtistID  int      `json:"artist_id"`
	Title     string   `json:"title"`
	AlbumID   *int     `json:"album_id,omitempty"`
	Mood      string   `json:"mood"`
	Source    string   `json:"source"`
	Genres    []string `json:"genres,omitempty"`
	SpotifyID string   `json:"spotify_id,omitempty"`
	ImagePath string   `json:"image_path,omitempty"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type SongRequest struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title" binding:"required"`
	ArtistID  int    `json:"artist_id" binding:"required"`
	AlbumID   *int   `json:"album_id,omitempty"`
	Mood      string `json:"mood" binding:"required,oneof=happy sad energetic calm angry relaxed"`
	Source    string `json:"source" binding:"oneof=manual spotify"`
	IDGenres  []int  `json:"genres,omitempty"`
	SpotifyID string `json:"spotify_id,omitempty"`
	ImagePath string `json:"image_path,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
