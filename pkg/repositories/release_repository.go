package repositories

import (
	"recoshelf-api/pkg/entities"

	"github.com/jmoiron/sqlx"
)

type ReleaseRepository interface {
	GetReleasesByUser(userID int) (*[]entities.Release, error)
}

type releaseRepository struct {
	DB *sqlx.DB
}

func NewReleaseRepo(db *sqlx.DB) ReleaseRepository {
	return &releaseRepository{
		DB: db,
	}
}

func (r *releaseRepository) GetReleasesByUser(userID int) (*[]entities.Release, error) {
	q := `
		SELECT releases.* 
		FROM releases 
		LEFT JOIN releases_users ON releases.id = releases_users.release_id 
		WHERE releases_users.user_id = ?
	`

	releases := []entities.Release{}
	err := r.DB.Select(&releases, q, userID)

	return &releases, err
}
