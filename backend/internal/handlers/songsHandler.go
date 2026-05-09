package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"music-mood-api/internal/models"
	"music-mood-api/pkg/errors"
)

type SongService interface {
	GetAll(f models.SongFilters) ([]models.Song, int, error)
	GetByID(id int) (*models.Song, error)
	Create(req models.SongRequest) (*models.Song, error)
	Update(id int, req models.SongRequest) (*models.Song, error)
	Delete(id int) error
	UpdateImage(id int, imagePath string) error
}

type SongHandler struct {
	service SongService
}

func NewSongHandler(service SongService) *SongHandler {
	return &SongHandler{service: service}
}

func (h *SongHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	limit, _ := strconv.Atoi(q.Get("limit"))

	filters := models.SongFilters{
		Page:   page,
		Limit:  limit,
		Search: q.Get("q"),
		Sort:   q.Get("sort"),
		Order:  q.Get("order"),
		Mood:   q.Get("mood"),
	}

	songs, total, err := h.service.GetAll(filters)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"data":  songs,
		"total": total,
		"page":  filters.Page,
		"limit": filters.Limit,
	})
}

func (h *SongHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid song ID")
		return
	}

	song, err := h.service.GetByID(id)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, song)
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

func (h *SongHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid song ID")
		return
	}

	var req models.SongRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	song, err := h.service.Update(id, req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, song)
}

func (h *SongHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid song ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		if err == sql.ErrNoRows {
			writeError(w, http.StatusNotFound, "song not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}
