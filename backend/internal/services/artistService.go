package services

import (
	"fmt"
	"music-mood-api/internal/models"
	"music-mood-api/pkg/errors"
)

type ArtistRepository interface {
	GetAll() ([]models.Artist, error)
	GetByID(id int) (*models.Artist, error)
	Create(req models.ArtistRequest) (*models.Artist, error)
	Update(id int, req models.ArtistRequest) (*models.Artist, error)
	Delete(id int) error
}

type ArtistService struct {
	repo ArtistRepository
}

func NewArtistService(repo ArtistRepository) *ArtistService {
	return &ArtistService{repo: repo}
}

func (s *ArtistService) GetAll() ([]models.Artist, error) {
	return s.repo.GetAll()
}

func (s *ArtistService) GetByID(id int) (*models.Artist, error) {
	artist, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.Internal("error retrieving artist")
	}
	if artist == nil {
		return nil, errors.NotFound(fmt.Sprintf("artist with ID %d not found", id))
	}
	return artist, nil
}

func (s *ArtistService) Create(req models.ArtistRequest) (*models.Artist, error) {
	if req.Name == "" {
		return nil, errors.BadRequest("artist name is required")
	}
	return s.repo.Create(req)
}

func (s *ArtistService) Update(id int, req models.ArtistRequest) (*models.Artist, error) {
	if req.Name == "" {
		return nil, errors.BadRequest("artist name is required")
	}
	artist, err := s.repo.Update(id, req)
	if err != nil {
		return nil, errors.Internal("error updating artist")
	}
	if artist == nil {
		return nil, errors.NotFound(fmt.Sprintf("artist with ID %d not found", id))
	}
	return artist, nil
}

func (s *ArtistService) Delete(id int) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return errors.Internal("error retrieving artist")
	}
	if existing == nil {
		return errors.NotFound(fmt.Sprintf("artist with ID %d not found", id))
	}
	return s.repo.Delete(id)
}
