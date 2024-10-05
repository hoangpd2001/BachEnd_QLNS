package reqUser

import "database/sql"

type ReqUser struct {
	Ten            string        `validate:"required"`
	Dem            string        `validate:"required"`
	Ho             string        `validate:"required"`
	Email          string        `validate:"required"`
	NgaySinh       string        `validate:"required"`
	IDLoaiNhanVien sql.NullInt64 `validate:"required"`
	IDCapBac       sql.NullInt64 `validate:"required"`
	GioiTinh       string        `validate:""`
	SDT            string        `validate:""`
	DiaChi         string        `validate:""`
	CCCD           string        `validate:""`
}
