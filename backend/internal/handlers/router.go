package handlers

import (
	"database/sql"
	"music-mood-api/internal/db"
	"music-mood-api/internal/services"
	"net/http"
)

func NewRouter(database *sql.DB) http.Handler {
	mux := http.NewServeMux()

	repo := db.NewSongRepository(database)
	songService := services.NewSongService(repo)
	songHandler := NewSongHandler(songService)

	mux.HandleFunc("GET /songs", songHandler.GetAll)

	return mux
}
