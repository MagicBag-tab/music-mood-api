package models

type Album struct {
	ID          int      `json:"id,omitempty"`
	ArtistID    int      `json:"artist_id"`
	Title       string   `json:"title"`
	Songs       []string `json:"songs,omitempty"`
	ReleaseDate string   `json:"release_date,omitempty"`
	CoverPath   string   `json:"cover_path,omitempty"`
	CreatedAt   string   `json:"created_at,omitempty"`
}

type AlbumRequest struct {
	ID          int    `json:"id,omitempty"`
	ArtistID    int    `json:"artist_id"`
	Title       string `json:"title"`
	Songs       []int  `json:"songs,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
	CoverPath   string `json:"cover_path,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}
