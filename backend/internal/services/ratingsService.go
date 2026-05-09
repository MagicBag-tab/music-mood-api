package services

import (
	"music-mood-api/internal/models"
	"music-mood-api/pkg/errors"
)

type RatingRepository interface {
	Create(songID int, req models.RatingRequest) (*models.Rating, error)
	GetBySongID(songID int) ([]models.Rating, error)
	GetAverageBySongID(songID int) (float64, error)
}

type RatingSongRepository interface {
	GetByID(id int) (*models.Song, error)
}

type RatingService struct {
	ratingRepo RatingRepository
	songRepo   RatingSongRepository
}

func NewRatingService(ratingRepo RatingRepository, songRepo RatingSongRepository) *RatingService {
	return &RatingService{ratingRepo: ratingRepo, songRepo: songRepo}
}

func (s *RatingService) Create(songID int, req models.RatingRequest) (*models.Rating, error) {
	if req.Score < 1 || req.Score > 5 {
		return nil, errors.BadRequest("score must be between 1 and 5")
	}

	song, err := s.songRepo.GetByID(songID)
	if err != nil {
		return nil, errors.Internal("error verifying song")
	}
	if song == nil {
		return nil, errors.NotFound("song not found")
	}

	return s.ratingRepo.Create(songID, req)
}

type RatingSummary struct {
	Ratings []models.Rating `json:"ratings"`
	Average float64         `json:"average"`
	Total   int             `json:"total"`
}

func (s *RatingService) GetBySongID(songID int) (*RatingSummary, error) {
	song, err := s.songRepo.GetByID(songID)
	if err != nil {
		return nil, errors.Internal("error verifying song")
	}
	if song == nil {
		return nil, errors.NotFound("song not found")
	}

	ratings, err := s.ratingRepo.GetBySongID(songID)
	if err != nil {
		return nil, errors.Internal("error retrieving ratings")
	}

	avg, err := s.ratingRepo.GetAverageBySongID(songID)
	if err != nil {
		return nil, errors.Internal("error calculating average")
	}

	if ratings == nil {
		ratings = []models.Rating{}
	}

	return &RatingSummary{
		Ratings: ratings,
		Average: avg,
		Total:   len(ratings),
	}, nil
}
