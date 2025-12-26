package repositories

import (
	"errors"

	"recoshelf-api/pkg/entities"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ReleaseRepository interface {
	GetReleaseBySource(source string, sourceReleaseID int64) (*entities.Release, error)
	GetReleasesByUser(userID int64) (*[]entities.Release, error)
	CreateRelease(release entities.Release) (int64, error)
	CreateReleaseUserRelation(userID int64, releaseID int64) error
	DeleteReleaseUserRelation(userID int64, releaseID int64) error
}

type releaseRepository struct {
	DB *sqlx.DB
}

func NewReleaseRepo(db *sqlx.DB) ReleaseRepository {
	return &releaseRepository{
		DB: db,
	}
}

func (r *releaseRepository) GetReleaseBySource(source string, sourceReleaseID int64) (*entities.Release, error) {
	q := `
		SELECT * FROM releases WHERE source = ? AND source_release_id = ?
	`

	release := entities.Release{}
	err := r.DB.Get(&release, q, source, sourceReleaseID)

	return &release, err
}

func (r *releaseRepository) GetReleasesByUser(userID int64) (*[]entities.Release, error) {
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

func (r *releaseRepository) CreateRelease(release entities.Release) (int64, error) {
	q := `
		INSERT INTO releases (source, source_release_id, title, country, genres, release_year, tracklist, images, barcode, fetched_at)
		VALUES (:source, :source_release_id, :title, :country, :genres, :release_year, :tracklist, :images, :barcode, :fetched_at)
	`

	result, err := r.DB.NamedExec(q, release)
	if err != nil {
		return 0, err
	}

	releaseID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return releaseID, nil
}

func (r *releaseRepository) CreateReleaseUserRelation(userID int64, releaseID int64) error {
	q := `
		INSERT INTO releases_users (release_id, user_id) VALUES (:releaseID, :userID)
	`

	_, err := r.DB.NamedExec(q, map[string]interface{}{
		"userID":    userID,
		"releaseID": releaseID,
	})

	if err != nil {
		var mysqlErr *mysql.MySQLError
		// Check if duplicate
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil
		}
	}

	return err
}

func (r *releaseRepository) DeleteReleaseUserRelation(userID int64, releaseID int64) error {
	q := `
		DELETE FROM releases_users WHERE release_id = :releaseID AND user_id = :userID
	`

	_, err := r.DB.NamedExec(q, map[string]interface{}{
		"userID":    userID,
		"releaseID": releaseID,
	})

	return err
}
