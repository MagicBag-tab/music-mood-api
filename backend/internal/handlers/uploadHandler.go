package handlers

import (
	"fmt"
	"io"
	"music-mood-api/pkg/errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const maxUploadSize = 1 << 20

func (h *SongHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	songID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid song ID")
		return
	}

	song, err := h.service.GetByID(songID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error checking song")
		return
	}
	if song == nil {
		writeError(w, http.StatusNotFound, "song not found")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		writeError(w, http.StatusBadRequest, "image too large, max 1MB")
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		writeError(w, http.StatusBadRequest, "image field is required")
		return
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	if !allowed[ext] {
		writeError(w, http.StatusBadRequest, "only jpg, jpeg, png, webp allowed")
		return
	}

	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		writeError(w, http.StatusInternalServerError, "could not create uploads dir")
		return
	}

	filename := fmt.Sprintf("song_%d_%d%s", songID, time.Now().Unix(), ext)
	destPath := filepath.Join("uploads", filename)

	dest, err := os.Create(destPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not save image")
		return
	}
	defer dest.Close()

	if _, err := io.Copy(dest, file); err != nil {
		writeError(w, http.StatusInternalServerError, "could not write image")
		return
	}

	imagePath := "/uploads/" + filename
	if err := h.service.UpdateImage(songID, imagePath); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			writeError(w, appErr.Code, appErr.Message)
			return
		}
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"image_path": imagePath,
	})
}
