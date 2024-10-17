package reqUser

type ReqEducation struct {
	IDNhanVien   int    `validate:""`
	Truong       string `validate:""`
	BangCap      string `validate:""`
	CapHoc       string `validate:""`
	NamTotNghiep string `validate:""`
}
