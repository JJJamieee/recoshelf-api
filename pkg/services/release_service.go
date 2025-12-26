package services

import (
	"database/sql"
	"errors"

	"recoshelf-api/pkg/entities"
	"recoshelf-api/pkg/repositories"
)

type ReleaseService interface {
	GetUserReleases(userID int64) (*[]entities.Release, error)
	CreateUserRelease(userID int64, release entities.Release) error
}

type releaseService struct {
	releaseRepository repositories.ReleaseRepository
}

func NewReleaseService(rr repositories.ReleaseRepository) ReleaseService {
	return &releaseService{
		releaseRepository: rr,
	}
}

func (s *releaseService) GetUserReleases(userID int64) (*[]entities.Release, error) {
	return s.releaseRepository.GetReleasesByUser(userID)
}

func (s *releaseService) CreateUserRelease(userID int64, release entities.Release) error {
	existRelease, err := s.releaseRepository.GetReleaseBySource(release.Source, release.SourceReleaseID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	var releaseID int64
	if errors.Is(err, sql.ErrNoRows) || existRelease == nil {
		releaseID, err = s.releaseRepository.CreateRelease(release)
		if err != nil {
			return err
		}
	} else {
		releaseID = existRelease.ID
	}

	err = s.releaseRepository.CreateReleaseUserRelation(userID, releaseID)

	return err
}
