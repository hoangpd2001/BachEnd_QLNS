package reqUser



type ReqUser struct {
	Ten            string        `validate:"required"`
	Dem            string        `validate:"required"`
	Ho             string        `validate:"required"`
	Email          string        `validate:"required"`
	NgaySinh       string        `validate:"required"`
	IDLoaiNhanVien int `validate:"required"`
	IDCapBac       int `validate:"required"`
	GioiTinh       string        `validate:""`
	SDT            string        `validate:""`
	DiaChi         string        `validate:""`
	CCCD           string        `validate:""`
}
