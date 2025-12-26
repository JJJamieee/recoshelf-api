package requests

import (
	"time"

	"recoshelf-api/pkg/entities"
)

type ReleaseRequest struct {
	Source          string         `json:"source" validate:"required"`
	SourceReleaseID int64          `json:"sourceReleaseId" validate:"required"`
	Title           string         `json:"title" validate:"required"`
	Country         string         `json:"country" validate:"required"`
	Genres          []string       `json:"genres" validate:"required"`
	ReleaseYear     uint           `json:"releaseYear" validate:"required"`
	Tracklist       []TrackRequest `json:"tracklist" validate:"required"`
	Images          *string        `json:"images"`
	Barcode         string         `json:"barcode" validate:"required"`
	FetchedAt       time.Time      `json:"fetchedAt" validate:"required"`
}

type TrackRequest struct {
	Duration string `json:"duration" validate:"required"`
	Title    string `json:"title" validate:"required"`
}

func (r ReleaseRequest) ToEntity() entities.Release {
	tracklist := make(entities.TrackList, 0, len(r.Tracklist))
	for _, track := range r.Tracklist {
		tracklist = append(tracklist, entities.Track{
			Duration: track.Duration,
			Title:    track.Title,
		})
	}

	return entities.Release{
		Source:          r.Source,
		SourceReleaseID: r.SourceReleaseID,
		Title:           r.Title,
		Country:         r.Country,
		Genres:          entities.StringSlice(r.Genres),
		ReleaseYear:     r.ReleaseYear,
		Tracklist:       tracklist,
		Images:          r.Images,
		Barcode:         r.Barcode,
		FetchedAt:       r.FetchedAt.UTC(),
	}
}
