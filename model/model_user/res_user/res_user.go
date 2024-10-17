package resuser

import (
	
	"time"
)

type ResUser struct {
	ID             int       `db:"ID"             json:"ID,omitempty"`
	Ten            string    `db:"Ten"            json:"Ten,omitempty"`
	Dem            string    `db:"Dem"            json:"Dem,omitempty"`
	Ho             string    `db:"Ho"             json:"Ho,omitempty"`
	Email          string    `db:"Email"          json:"Email,omitempty"`
	NgaySinh       time.Time `db:"NgaySinh"       json:"NgaySinh,omitempty"`
	GioiTinh       string    `db:"GioiTinh"       json:"GioiTinh,omitempty"`
	SDT            string    `db:"SDT"            json:"SDT,omitempty"`
	DiaChi         string    `db:"DiaChi"         json:"DiaChi,omitempty"`
	CCCD           string    `db:"CCCD"           json:"CCCD,omitempty"`

	LoaiNhanVien  string    `db:"LoaiNhanVien"   json:"LoaiNhanVien,omitempty"`
	IDLoaiNhanVien int      `db:"IDLoaiNhanVien" json:"IDLoaiNhanVien,omitempty"`
	
	IDCapBac       int      `db:"IDCapBac"       json:"IDCapBac,omitempty"`
	TenCapBac     string    `db:"TenCapBac"      json:"TenCapBac,omitempty"`
	
	IDChiNhanh     int      `db:"IDChiNhanh"     json:"IDChiNhanh,omitempty"`
	ChiNhanh      string    `db:"ChiNhanh"       json:"ChiNhanh,omitempty"`
	
	IDPhongBan     int      `db:"IDPhongBan"     json:"IDPhongBan,omitempty"`
	TenPhongBan   string       `db:"TenPhongBan"    json:"TenPhongBan,omitempty"`
	
	IDChucDanh     int      `db:"IDChucDanh"     json:"IDChucDanh,omitempty"`
	TenChucDanh   string    `db:"TenChucDanh"    json:"TenChucDanh,omitempty"`
	
	
	NgayBatDau     time.Time `db:"NgayBatDau"     json:"NgayBatDau,omitempty"`
	NgayKetThuc    time.Time `db:"NgayKetThuc"    json:"NgayKetThuc,omitempty"`

	TenNguoiThan   string `db:"TenNguoiThan"     json:"TenNguoiThan,omitempty"`
	SDTNguoiThan   string `db:"SDTNguoiThan"     json:"SDTNguoiThan,omitempty"`
	DiaChiNguoiThan string `db:"DiaChiNguoiThan"  json:"DiaChiNguoiThan,omitempty"`
	QuanHe         string `db:"QuanHe"           json:"QuanHe,omitempty"`
}
