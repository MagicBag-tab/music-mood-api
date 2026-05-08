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
	songService := services.NewSongService(songrepo, artistrepo)
	songHandler := NewSongHandler(songService)

	mux.HandleFunc("GET /songs", songHandler.GetAll)
	mux.HandleFunc("POST /songs", songHandler.Create)

	return mux
}
