package reqUser

import "time"

type ReqUser struct {
	Ten          string    `validate:"required"`
	Dem          string    `validate:"required"`
	Ho           string    `validate:"required"`
	Email        string    `validate:"required"`
	NgaySinh     time.Time `validate:"required"`
	LoaiNhanVien int       `validate:"required"`
	CapBac       int       `validate:"required"`
	ChiNhanh     int       `validate:"required"`
	GioiTinh     bool      `validate:""`
	SDT          string    `validate:""`
	DiaChi       string    `validate:""`
	CCCD         string    `validate:""`
}
