package entities

import (
	"time"
)

type Release struct {
	ID              int64       `json:"id" db:"id"`
	Source          string      `json:"-" db:"source"`
	SourceReleaseID int64       `json:"-" db:"source_release_id"`
	Title           string      `json:"title" db:"title"`
	Country         string      `json:"country" db:"country"`
	Genres          StringSlice `json:"genres" db:"genres"`
	ReleaseYear     uint        `json:"releaseYear" db:"release_year"`
	Tracklist       TrackList   `json:"tracklist" db:"tracklist"`
	Images          *string     `json:"images" db:"images"`
	Barcode         string      `json:"barcode" db:"barcode"`
	FetchedAt       time.Time   `json:"-" db:"fetched_at"`
	CreatedAt       time.Time   `json:"-" db:"created_at"`
	UpdatedAt       time.Time   `json:"-" db:"updated_at"`
}
