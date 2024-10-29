package modelrole

type ReqUserRole struct {
	IDChucDanh int `validate:""`
	IDVaiTro   int `validate:""`
	Xem        bool `validate:""`
	Them       bool `validate:""`
	Sua        bool `validate:""`
	Xoa        bool `validate:""`
}

type ReqRole struct {
	ID int `validate:"required"`
	Ten   string `validate:"required"`
}

