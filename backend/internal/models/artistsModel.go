package models

type Artist struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Country   string `json:"country,omitempty"`
	ImagePath string `json:"image_path,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type ArtistRequest struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Country   string `json:"country,omitempty"`
	ImagePath string `json:"image_path,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
