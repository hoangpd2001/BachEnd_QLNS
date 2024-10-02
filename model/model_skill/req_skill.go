package modelskill

import "time"

type ReqSkill struct {
	ID        int    `db:"ID"  json:"ID,omitempty"`
	TenKyNang string `db:"TenKyNang"          json:"TenKyNang,omitempty"`
	MoTa      string `db:"MoTa"          json:"MoTa,omitempty"`
}

type Req_user_skill struct {
	IDNhanVien  int       `db:"IDNhanVien"  json:"IDNhanVien,omitempty"`
	IDKyNang    int       `db:"IDKyNang"  json:"IDKyNang,omitempty"`
	MucDo       int       `db:"MucDo"          json:"MucDo,omitempty"`
	NgayDanhGia time.Time `db:"NgayDanhGia"          json:"NgayDanhGia,omitempty"`
}
