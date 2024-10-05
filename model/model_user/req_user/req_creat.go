package reqUser

type ReqCreate struct {

	Ten          string    `validate:"required"`
	Dem          string    `validate:"required"`
	Ho           string    `validate:"required"`
	Email        string    `validate:"required"`
	Pass         string    `validate:"required"`
	LoaiNhanVien int       `validate:"required"`
	CapBac       int       `validate:"required"`
	ChiNhanh     int       `validate:"required"`

	GioiTinh        bool   `validate:""`
	SDT             string `validate:""`
	EmailCaNhan     string `validate:""`
	DiaChiThuongTru string `validate:""`
	DiaChiTamTru    string `validate:""`
	CCCD            string `validate:""`

	//NguoiThan  `validate:""`
	TenNguoiThan    string `validate:""`
	SDTNguoiThan    string `validate:""`
	DiaChiNguoiThan string `validate:""`
	QuanHe          string `validate:""`
}

