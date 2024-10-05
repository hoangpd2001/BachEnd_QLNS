package modelgrup

type ReqGrup struct {
	TenNhom string `validate:"required"`
}

type ReqUserGrup struct {
	IDNhanVien int `validate:""`
	IDNhom     int `validate:""`
}
