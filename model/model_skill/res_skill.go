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
	MucDo       int       `db:"MucDo"            json:"MucDo,omitempty"`
	Ten            string    `db:"Ten"            json:"Ten,omitempty"`
	Dem            string    `db:"Dem"            json:"Dem,omitempty"`
	Ho             string    `db:"Ho"             json:"Ho,omitempty"`
	NgayDanhGia time.Time `db:"NgayDanhGia"      json:"NgayDanhGia,omitempty"`
	ChucDanh    string    `db:"ChucDanh"         json:"ChucDanh,omitempty"`
	TenKyNang   string    `db:"TenKyNang"        json:"TenKyNang,omitempty"`
	MoTa        string    `db:"MoTa"             json:"MoTa,omitempty"`
}
