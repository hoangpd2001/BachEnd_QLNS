package utils

import (
	"fmt"
	"time"
)

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	// Parse chuỗi chỉ chứa ngày từ JSON ("YYYY-MM-DD")
	t, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	// Chuyển thành chuỗi chỉ chứa ngày để trả về JSON ("YYYY-MM-DD")
	formatted := fmt.Sprintf("\"%s\"", cd.Time.Format("2006-01-02"))
	return []byte(formatted), nil
}
