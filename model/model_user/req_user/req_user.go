package reqUser

type ReqUser struct {
	Ten          string    `validate:"required"`
	Dem          string    `validate:"required"`
	Ho           string    `validate:"required"`
	Email        string    `validate:"required"`
	LoaiNhanVien int       `validate:"required"`
	CapBac       int       `validate:"required"`
	ChiNhanh     int       `validate:"required"`
}
