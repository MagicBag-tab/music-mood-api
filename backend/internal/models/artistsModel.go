package models

import "time"

type Artist struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	ImagePath string    `json:"image_path,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type ArtistRequest struct {
	Name      string `json:"name"`
	Country   string `json:"country"`
	ImagePath string `json:"image_path"`
}
