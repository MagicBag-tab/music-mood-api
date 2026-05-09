package services

import (
	"music-mood-api/internal/models"
	"music-mood-api/pkg/errors"
)

type ReportRepository interface {
	GetMoodDistribution() ([]models.MoodDistribution, error)
	GetTopRatedSongs(limit int) ([]models.TopRatedSong, error)
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetMoodDistribution() ([]models.MoodDistribution, error) {
	data, err := s.repo.GetMoodDistribution()
	if err != nil {
		return nil, errors.Internal("error generating mood distribution report")
	}
	return data, nil
}

func (s *ReportService) GetTopRatedSongs(limit int) ([]models.TopRatedSong, error) {
	if limit <= 0 {
		limit = 5 // Default value
	}

	data, err := s.repo.GetTopRatedSongs(limit)
	if err != nil {
		return nil, errors.Internal("error generating top rated songs report")
	}
	return data, nil
}
