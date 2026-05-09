package handlers

import (
	"database/sql"
	"music-mood-api/internal/db"
	"music-mood-api/internal/services"
	"net/http"
)

func NewRouter(database *sql.DB) http.Handler {
	mux := http.NewServeMux()

	songrepo := db.NewSongRepository(database)
	artistrepo := db.NewArtistRepository(database)
	albumrepo := db.NewAlbumRepository(database)
	ratingrepo := db.NewRatingRepository(database)
	reportrepo := db.NewReportRepository(database)

	songService := services.NewSongService(songrepo, artistrepo)
	songHandler := NewSongHandler(songService)
	artistService := services.NewArtistService(artistrepo)
	artistHandler := NewArtistHandler(artistService)
	albumService := services.NewAlbumService(albumrepo, artistrepo)
	albumHandler := NewAlbumHandler(albumService)
	ratingService := services.NewRatingService(ratingrepo, songrepo)
	ratingHandler := NewRatingHandler(ratingService)
	reportService := services.NewReportService(reportrepo)
	reportHandler := NewReportHandler(reportService)

	mux.HandleFunc("GET /songs", songHandler.GetAll)
	mux.HandleFunc("POST /songs", songHandler.Create)
	mux.HandleFunc("GET /songs/{id}", songHandler.GetByID)
	mux.HandleFunc("PUT /songs/{id}", songHandler.Update)
	mux.HandleFunc("DELETE /songs/{id}", songHandler.Delete)
	mux.HandleFunc("POST /songs/{id}/rating", ratingHandler.Create)
	mux.HandleFunc("GET /songs/{id}/rating", ratingHandler.GetBySongID)

	mux.HandleFunc("GET /artists", artistHandler.GetAll)
	mux.HandleFunc("POST /artists", artistHandler.Create)
	mux.HandleFunc("GET /artists/{id}", artistHandler.GetByID)
	mux.HandleFunc("PUT /artists/{id}", artistHandler.Update)
	mux.HandleFunc("DELETE /artists/{id}", artistHandler.Delete)

	mux.HandleFunc("GET /albums", albumHandler.GetAll)
	mux.HandleFunc("POST /albums", albumHandler.Create)
	mux.HandleFunc("GET /albums/{id}", albumHandler.GetByID)
	mux.HandleFunc("PUT /albums/{id}", albumHandler.Update)
	mux.HandleFunc("DELETE /albums/{id}", albumHandler.Delete)

	mux.HandleFunc("GET /reports/moods", reportHandler.GetMoodDistribution)
	mux.HandleFunc("GET /reports/top-rated", reportHandler.GetTopRatedSongs)

	return mux
}
