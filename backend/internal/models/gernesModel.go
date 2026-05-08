package models

type Genre struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type GenreRequest struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}
