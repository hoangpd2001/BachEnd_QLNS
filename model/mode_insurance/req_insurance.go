package modeinsurance


type ReqInsurance struct {
	TenBaoHiem string `validate:"required"`
	NhaCungCap      string `validate:""`
	NoiDangKi      string `validate:""`
	TyLePhi      int `validate:""`
}

type ReqUserInsurance struct {
	IDNhanVien  int       `validate:""`
	IDBaoHiem    int       `validate:""`
	NgayDong       string       `validate:"required"`
	NgayHetHan string `validate:"required"`
}
