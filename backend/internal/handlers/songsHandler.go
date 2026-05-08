package handlers

import (
	"encoding/json"
	"music-mood-api/internal/models"
	"music-mood-api/pkg/errors"
	"net/http"
)

type SongService interface {
	GetAll() ([]models.Song, error)
	Create(req models.SongRequest) (*models.Song, error)
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

func (h *SongHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.SongRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	song, err := h.service.Create(req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusCreated, song)
}
