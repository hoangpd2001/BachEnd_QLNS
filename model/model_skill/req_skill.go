package modelskill


type ReqSkill struct {
	TenKyNang string `validate:"required"`
	MoTa      string `validate:""`
}

type ReqUserSkill struct {
	IDNhanVien  int       `validate:""`
	IDKyNang    int       `validate:""`
	MucDo       int       `validate:"required"`
	NgayDanhGia string `validate:"required"`
}
