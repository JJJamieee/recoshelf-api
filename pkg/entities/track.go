package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Track struct {
	Duration string
	Title    string
}

type TrackList []Track

func (t *TrackList) Scan(value any) error {
	if value == nil {
		*t = TrackList{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot scan %T into TrackList", value)
	}

	return json.Unmarshal(bytes, t)
}

func (t TrackList) Value() (driver.Value, error) {
	if len(t) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(t)
}
