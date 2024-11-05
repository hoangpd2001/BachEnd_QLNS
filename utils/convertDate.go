package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	var dateString string

	if err := json.Unmarshal(data, &dateString); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return err
	}

	cd.Time = t
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", cd.Time.Format("2006-01-02"))
	return []byte(formatted), nil
}
