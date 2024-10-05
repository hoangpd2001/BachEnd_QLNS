package modeltitle

import "time"

type ResTitle struct {
	ID          int    `db:"ID"                   json:"ID,omitempty"`
	TenChucDanh string `db:"TenChucDanh"          json:"TenChucDanh,omitempty"`
}
type ResUserTitle struct {
	IDNhanVien  int       `db:"IDNhanVien"        json:"IDNhanVien,omitempty"`
	IDChucDanh  int       `db:"IDChucDanh"        json:"IDChucDanh,omitempty"`
	IDPhongBan  int       `db:"IDPhongBan"        json:"IDPhongBan,omitempty"`
	Ten         string    `db:"Ten"               json:"Ten,omitempty"`
	Dem         string    `db:"Dem"               json:"Dem,omitempty"`
	Ho          string    `db:"Ho"                json:"Ho,omitempty"`
	TenPhongBan string    `db:"TenPhongBan"       json:"TenPhongBan,omitempty"`
	ChiNhanh    string    `db:"ChiNhanh"          json:"ChiNhanh,omitempty"`
	ID          int       `db:"ID"                json:"ID,omitempty"`
	NgayBatDau  time.Time `db:"NgayBatDau"        json:"NgayBatDau,omitempty"`
	NgayKetThuc time.Time `db:"NgayKetThuc"       json:"NgayKetThuc,omitempty"`
	TenChucDanh string `db:"TenChucDanh"          json:"TenChucDanh,omitempty"`

}
