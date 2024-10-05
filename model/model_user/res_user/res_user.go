package resuser

import "time"

type ResUser struct {
	ID           int       `db:"ID"             json:"ID,omitempty"`
	Ten          string    `db:"Ten"            json:"Ten,omitempty"`
	Dem          string    `db:"Dem"            json:"Dem,omitempty"`
	Ho           string    `db:"Ho"             json:"Ho,omitempty"`
	Email        string    `db:"Email"          json:"Email,omitempty"`
	LoaiNhanVien int       `db:"LoaiNhanVien"   json:"LoaiNhanVien,omitempty"`
	CapBac       int       `db:"CapBac"         json:"CapBac,omitempty"`
	ChiNhanh     int       `db:"ChiNhanh"       json:"ChiNhanh,omitempty"`
	NgaySinh     time.Time `db:"NgaySinh"    json:"NgaySinh,omitempty"`
	GioiTinh     string    `db:"GioiTinh"    json:"GioiTinh,omitempty"`
	SDT          string    `db:"SDT"    json:"SDT,omitempty"`
	DiaChi       string    `db:"DiaChi"    json:"DiaChi,omitempty"`
	CCCD         string    `db:"CCCD"    json:"CCCD,omitempty"`
	NgayBatDau   time.Time `db:"NgayBatDau"     json:"NgayBatDau,omitempty"`
	NgayKetThuc  time.Time `db:"NgayKetThuc"    json:"NgayKetThuc,omitempty"`
}
