package handlers

import (
	"music-mood-api/internal/services"
	"music-mood-api/pkg/errors"
	"net/http"
	"strconv"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetMoodDistribution(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetMoodDistribution()
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, data)
}

func (h *ReportHandler) GetTopRatedSongs(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit := 5
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	data, err := h.service.GetTopRatedSongs(limit)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, data)
}
