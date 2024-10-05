package modeltitle


type ReqTitle struct {
	TenChucDanh string `validate:"required"`
}

type ReqUserTitle struct {
	IDNhanVien  int       `validate:"required"`
	IDChucDanh  int       `validate:"required"`
	IDPhongBan  int       `validate:"required"`
	NgayBatDau  string `validate:"required"`
	NgayKetThuc string `validate:"required"`
}
