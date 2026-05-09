package models

type MoodDistribution struct {
	Mood  string `json:"mood"`
	Count int    `json:"count"`
}

type TopRatedSong struct {
	SongID        int     `json:"song_id"`
	Title         string  `json:"title"`
	ArtistName    string  `json:"artist_name"`
	AverageRating float64 `json:"average_rating"`
	RatingsCount  int     `json:"ratings_count"`
}
