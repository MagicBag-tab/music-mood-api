package services

import (
	"fmt"
	"music-mood-api/internal/models"
	"music-mood-api/pkg/errors"
)

type SongRepository interface {
	GetAll() ([]models.Song, error)
	Create(req models.SongRequest) (*models.Song, error)
}

type ArtistRepository interface {
	GetByID(id int) (*models.Artist, error)
}

type SongService struct {
	songrepo   SongRepository
	artistrepo ArtistRepository
}

func NewSongService(songrepo SongRepository, artistrepo ArtistRepository) *SongService {
	return &SongService{songrepo: songrepo, artistrepo: artistrepo}
}

func (s *SongService) GetAll() ([]models.Song, error) {
	return s.songrepo.GetAll()
}

func (s *SongService) Create(req models.SongRequest) (*models.Song, error) {
	if req.Title == "" {
		return nil, errors.BadRequest("title is required")
	}

	artist, err := s.artistrepo.GetByID(req.ArtistID)
	if err != nil {
		return nil, errors.Internal("error al verificar el artista")
	}
	if artist == nil {
		return nil, errors.NotFound(fmt.Sprintf("el artista con ID %d no existe", req.ArtistID))
	}

	validMoods := map[string]bool{
		"happy": true, "sad": true, "energetic": true,
		"calm": true, "angry": true, "relaxed": true,
	}
	if !validMoods[req.Mood] {
		return nil, errors.BadRequest(fmt.Sprintf("invalid mood: %s", req.Mood))
	}

	validSources := map[string]bool{
		"manual": true, "spotify": true,
	}
	if !validSources[req.Source] {
		return nil, errors.BadRequest(fmt.Sprintf("invalid source: %s", req.Source))
	}

	return s.songrepo.Create(req)
}
