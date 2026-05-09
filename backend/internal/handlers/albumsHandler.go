package handlers

import (
	"encoding/json"
	"music-mood-api/internal/models"
	"music-mood-api/pkg/errors"
	"net/http"
	"strconv"
)

type AlbumService interface {
	GetAll() ([]models.Album, error)
	GetByID(id int) (*models.Album, error)
	Create(req models.AlbumRequest) (*models.Album, error)
	Update(id int, req models.AlbumRequest) (*models.Album, error)
	Delete(id int) error
}

type AlbumHandler struct {
	service AlbumService
}

func NewAlbumHandler(service AlbumService) *AlbumHandler {
	return &AlbumHandler{service: service}
}

func (h *AlbumHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	albums, err := h.service.GetAll()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, albums)
}

func (h *AlbumHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid album ID")
		return
	}
	album, err := h.service.GetByID(id)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, album)
}

func (h *AlbumHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.AlbumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body")
		return
	}
	album, err := h.service.Create(req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, album)
}

func (h *AlbumHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid album ID")
		return
	}
	var req models.AlbumRequest
	json.NewDecoder(r.Body).Decode(&req)
	album, err := h.service.Update(id, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, album)
}

func (h *AlbumHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid album ID")
		return
	}
	if err := h.service.Delete(id); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}
