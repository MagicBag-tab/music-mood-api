package services

import "music-mood-api/internal/models"

type SongRepository interface {
	GetAll() ([]models.Song, error)
}

type SongService struct {
	repo SongRepository
}

func NewSongService(repo SongRepository) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) GetAll() ([]models.Song, error) {
	return s.repo.GetAll()
}
