package reqUser

type ReqRelative struct {
	TenNguoiThan    string `validate:"required"`
	SDTNguoiThan    string `validate:"required"`
	DiaChiNguoiThan string `validate:"required"`
	QuanHe          string `validate:"required"`
}
