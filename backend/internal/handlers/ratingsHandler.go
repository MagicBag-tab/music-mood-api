package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"music-mood-api/internal/models"
	"music-mood-api/internal/services"
	"music-mood-api/pkg/errors"
)

type RatingHandler struct {
	service *services.RatingService
}

func NewRatingHandler(service *services.RatingService) *RatingHandler {
	return &RatingHandler{service: service}
}

func (h *RatingHandler) Create(w http.ResponseWriter, r *http.Request) {
	songID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid song ID")
		return
	}

	var req models.RatingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	rating, err := h.service.Create(songID, req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusCreated, rating)
}

func (h *RatingHandler) GetBySongID(w http.ResponseWriter, r *http.Request) {
	songID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid song ID")
		return
	}

	summary, err := h.service.GetBySongID(songID)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, summary)
}
