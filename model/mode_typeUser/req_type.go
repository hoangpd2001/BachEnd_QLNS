package modetypeuser

type ReqTypeUser struct {
	LoaiNhanVien      string `validate:"required"`
}