package resuser

type ResRelative struct {
	IDNhanVien      int    `db:"IDNhanVien"      json:""`
	TenNguoiThan    string `db:"TenNguoiThan"             json:"TenNguoiThan,omitempty"`
	SDTNguoiThan    string `db:"SDTNguoiThan"             json:"SDTNguoiThan,omitempty"`
	DiaChiNguoiThan string `db:"DiaChiNguoiThan"          json:"DiaChiNguoiThan,omitempty"`
	QuanHe          string `db:"QuanHe"          json:"QuanHe,omitempty"`
}
