package services

import (
	"fmt"
	"music-mood-api/internal/models"
	"music-mood-api/pkg/errors"
)

type AlbumRepository interface {
	GetAll() ([]models.Album, error)
	GetByID(id int) (*models.Album, error)
	Create(req models.AlbumRequest) (*models.Album, error)
	Update(id int, req models.AlbumRequest) (*models.Album, error)
	Delete(id int) error
}

type AlbumService struct {
	repo       AlbumRepository
	artistRepo ArtistRepository
}

func NewAlbumService(repo AlbumRepository, artistRepo ArtistRepository) *AlbumService {
	return &AlbumService{repo: repo, artistRepo: artistRepo}
}

func (s *AlbumService) GetAll() ([]models.Album, error) {
	return s.repo.GetAll()
}

func (s *AlbumService) GetByID(id int) (*models.Album, error) {
	album, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.Internal("error retrieving album")
	}
	if album == nil {
		return nil, errors.NotFound(fmt.Sprintf("album with ID %d not found", id))
	}
	return album, nil
}

func (s *AlbumService) Create(req models.AlbumRequest) (*models.Album, error) {
	if req.Title == "" || req.ArtistID == 0 {
		return nil, errors.BadRequest("title and artist_id are required")
	}

	// Validar que el artista existe
	artist, _ := s.artistRepo.GetByID(req.ArtistID)
	if artist == nil {
		return nil, errors.NotFound(fmt.Sprintf("artist with ID %d not found", req.ArtistID))
	}

	return s.repo.Create(req)
}

func (s *AlbumService) Update(id int, req models.AlbumRequest) (*models.Album, error) {
	if req.Title == "" {
		return nil, errors.BadRequest("title is required")
	}
	return s.repo.Update(id, req)
}

func (s *AlbumService) Delete(id int) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return errors.Internal("error retrieving album")
	}
	if existing == nil {
		return errors.NotFound(fmt.Sprintf("album with ID %d not found", id))
	}
	return s.repo.Delete(id)
}
