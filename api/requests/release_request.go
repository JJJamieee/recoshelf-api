package requests

import "recoshelf-api/pkg/entities"

type ReleaseRequest struct {
	Title       string         `json:"title" validate:"required"`
	Country     string         `json:"country" validate:"required"`
	Genres      []string       `json:"genres" validate:"required"`
	ReleaseYear uint           `json:"releaseYear" validate:"required"`
	Tracklist   []TrackRequest `json:"tracklist" validate:"required"`
	Images      *string        `json:"images"`
	Barcode     string         `json:"barcode" validate:"required"`
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
		Title:       r.Title,
		Country:     r.Country,
		Genres:      entities.StringSlice(r.Genres),
		ReleaseYear: r.ReleaseYear,
		Tracklist:   tracklist,
		Images:      r.Images,
		Barcode:     r.Barcode,
	}
}
