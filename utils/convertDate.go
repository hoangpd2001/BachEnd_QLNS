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
	// Chuỗi để lưu ngày sau khi bỏ dấu ngoặc kép
	var dateString string

	// Giải mã JSON và lưu vào biến dateString
	if err := json.Unmarshal(data, &dateString); err != nil {
		return err
	}

	// Parse chuỗi ngày ("YYYY-MM-DD")
	t, err := time.Parse("2006-01-02", dateString)
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
