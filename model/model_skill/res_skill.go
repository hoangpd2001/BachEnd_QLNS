package modelskill

type ResSkill struct {
	ID int
	TenKyNang string `validate:"required"`
	MoTa string `validate:""`
}


type Res_user_skill struct {
	IDNhanVien int
	IDKyNang int
	MucDo   int    `validate:"required"`
	NgayDanhGia       string `validate:"required"`
} 

