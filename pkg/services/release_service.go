package services

import (
	"recoshelf-api/pkg/entities"
	"recoshelf-api/pkg/repositories"
)

type ReleaseService interface {
	GetUserReleases(userID int) (*[]entities.Release, error)
}

type releaseService struct {
	releaseRepository repositories.ReleaseRepository
}

func NewReleaseService(rr repositories.ReleaseRepository) ReleaseService {
	return &releaseService{
		releaseRepository: rr,
	}
}

func (s *releaseService) GetUserReleases(userID int) (*[]entities.Release, error) {
	return s.releaseRepository.GetReleasesByUser(userID)
}
