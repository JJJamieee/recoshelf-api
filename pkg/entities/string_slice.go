package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type StringSlice []string

func (s *StringSlice) Scan(value any) error {
	if value == nil {
		*s = []string{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan StringSlice: %T", value)
	}

	return json.Unmarshal(bytes, s)
}

func (s StringSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}
