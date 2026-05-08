package handlers

import (
	"encoding/json"
	"music-mood-api/internal/models"
	"net/http"
)

type SongService interface {
	GetAll() ([]models.Song, error)
}

type SongHandler struct {
	service SongService
}

func NewSongHandler(service SongService) *SongHandler {
	return &SongHandler{service: service}
}

func (h *SongHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	songs, err := h.service.GetAll()
	if err != nil {
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	if songs == nil {
		songs = []models.Song{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}
