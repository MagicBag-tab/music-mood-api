package services

import (
	"fmt"
	"music-mood-api/internal/models"
	"music-mood-api/pkg/errors"
)

type SongRepository interface {
	GetAll(f models.SongFilters) ([]models.Song, int, error)
	GetByID(id int) (*models.Song, error)
	Create(req models.SongRequest) (*models.Song, error)
	Update(id int, req models.SongRequest) (*models.Song, error)
	Delete(id int) error
	UpdateImage(id int, imagePath string) error
}

type SongService struct {
	songrepo   SongRepository
	artistrepo ArtistRepository
}

func NewSongService(songrepo SongRepository, artistrepo ArtistRepository) *SongService {
	return &SongService{songrepo: songrepo, artistrepo: artistrepo}
}

func (s *SongService) GetAll(f models.SongFilters) ([]models.Song, int, error) {
	return s.songrepo.GetAll(f)
}

func (s *SongService) GetByID(id int) (*models.Song, error) {
	song, err := s.songrepo.GetByID(id)
	if err != nil {
		return nil, errors.Internal("error retrieving song")
	}
	if song == nil {
		return nil, errors.NotFound(fmt.Sprintf("song with ID %d not found", id))
	}
	return song, nil
}

func (s *SongService) Create(req models.SongRequest) (*models.Song, error) {
	if req.Title == "" {
		return nil, errors.BadRequest("title is required")
	}
	if req.ArtistID == 0 {
		return nil, errors.BadRequest("artist_id is required")
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

	artist, err := s.artistrepo.GetByID(req.ArtistID)
	if err != nil {
		return nil, errors.Internal("error verifying artist")
	}
	if artist == nil {
		return nil, errors.NotFound(fmt.Sprintf("artist with ID %d not found", req.ArtistID))
	}

	return s.songrepo.Create(req)
}

func (s *SongService) Update(id int, req models.SongRequest) (*models.Song, error) {
	if req.Title == "" {
		return nil, errors.BadRequest("title is required")
	}

	validMoods := map[string]bool{
		"happy": true, "sad": true, "energetic": true,
		"calm": true, "angry": true, "relaxed": true,
	}
	if !validMoods[req.Mood] {
		return nil, errors.BadRequest(fmt.Sprintf("invalid mood: %s", req.Mood))
	}

	existing, err := s.songrepo.GetByID(id)
	if err != nil {
		return nil, errors.Internal("error retrieving song")
	}
	if existing == nil {
		return nil, errors.NotFound(fmt.Sprintf("song with ID %d not found", id))
	}

	if req.ArtistID != 0 {
		artist, err := s.artistrepo.GetByID(req.ArtistID)
		if err != nil {
			return nil, errors.Internal("error verifying artist")
		}
		if artist == nil {
			return nil, errors.NotFound(fmt.Sprintf("artist with ID %d not found", req.ArtistID))
		}
	}

	return s.songrepo.Update(id, req)
}

func (s *SongService) Delete(id int) error {
	existing, err := s.songrepo.GetByID(id)
	if err != nil {
		return errors.Internal("error retrieving song")
	}
	if existing == nil {
		return errors.NotFound(fmt.Sprintf("song with ID %d not found", id))
	}
	return s.songrepo.Delete(id)
}

func (s *SongService) UpdateImage(id int, imagePath string) error {
	return s.songrepo.UpdateImage(id, imagePath)
}
