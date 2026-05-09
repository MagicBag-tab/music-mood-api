package handlers

import (
	"database/sql"
	"encoding/json"
	"music-mood-api/internal/models"
	"music-mood-api/pkg/errors"
	"net/http"
	"strconv"
)

type ArtistService interface {
	GetAll() ([]models.Artist, error)
	GetByID(id int) (*models.Artist, error)
	Create(req models.ArtistRequest) (*models.Artist, error)
	Update(id int, req models.ArtistRequest) (*models.Artist, error)
	Delete(id int) error
}

type ArtistHandler struct {
	service ArtistService
}

func NewArtistHandler(service ArtistService) *ArtistHandler {
	return &ArtistHandler{service: service}
}

func (h *ArtistHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	artists, err := h.service.GetAll()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if artists == nil {
		artists = []models.Artist{}
	}
	writeJSON(w, http.StatusOK, artists)
}

func (h *ArtistHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid artist ID")
		return
	}

	artist, err := h.service.GetByID(id)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, artist)
}

func (h *ArtistHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.ArtistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	artist, err := h.service.Create(req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusCreated, artist)
}

func (h *ArtistHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid artist ID")
		return
	}

	var req models.ArtistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	artist, err := h.service.Update(id, req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, artist)
}

func (h *ArtistHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid artist ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		if err == sql.ErrNoRows {
			writeError(w, http.StatusNotFound, "artist not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}
