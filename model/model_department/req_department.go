package modeldepartment

type Reqdepartment struct {
	IDChiNhanh  int       `validate:"required"`
	TenPhongBan string `validate:"required"`
}
