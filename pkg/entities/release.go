package entities

import (
	"time"
)

type Release struct {
	ID          int64       `json:"id" db:"id"`
	Title       string      `json:"title" db:"title"`
	Country     string      `json:"country" db:"country"`
	Genres      StringSlice `json:"genres" db:"genres"`
	ReleaseYear uint        `json:"releaseYear" db:"release_year"`
	Tracklist   TrackList   `json:"tracklist" db:"tracklist"`
	Images      *string     `json:"images" db:"images"`
	Barcode     string      `json:"barcode" db:"barcode"`
	CreatedAt   time.Time   `json:"-" db:"created_at"`
	UpdatedAt   time.Time   `json:"-" db:"updated_at"`
}
