package modelskill

import "time"

type ResSkill struct {
	ID        int    `db:"ID"                 json:"ID,omitempty"`
	TenKyNang string `db:"TenKyNang"          json:"TenKyNang,omitempty"`
	MoTa      string `db:"MoTa"               json:"MoTa,omitempty"`
}

type ResUserSkill struct {
	IDNhanVien  int       `db:"IDNhanVien"       json:"IDNhanVien,omitempty"`
	IDKyNang    int       `db:"IDKyNang"         json:"IDKyNang,omitempty"`
	IDKyNangMoi int       `db:"IDKyNangMoi"      json:"IDKyNangMoi,omitempty"`
	MucDo       int       `db:"MucDo"            json:"MucDo,omitempty"`
	NgayDanhGia time.Time `db:"NgayDanhGia"      json:"NgayDanhGia,omitempty"`
	HoTen       string    `db:"HoTen"            json:"HoTen,omitempty"`
	ChucDanh    string    `db:"ChucDanh"         json:"ChucDanh,omitempty"`
	TenKyNang   string    `db:"TenKyNang"        json:"TenKyNang,omitempty"`
	MoTa        string    `db:"MoTa"             json:"MoTa,omitempty"`
}
