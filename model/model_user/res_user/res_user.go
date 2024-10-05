package resuser

import (
	"database/sql"
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

	LoaiNhanVien   sql.NullString    `db:"LoaiNhanVien"   json:"LoaiNhanVien,omitempty"`
	IDLoaiNhanVien sql.NullInt64       `db:"IDLoaiNhanVien" json:"IDLoaiNhanVien,omitempty"`
	
	IDCapBac       sql.NullInt64       `db:"IDCapBac"       json:"IDCapBac,omitempty"`
	TenCapBac      sql.NullString    `db:"TenCapBac"      json:"TenCapBac,omitempty"`
	
	IDChiNhanh     sql.NullInt64       `db:"IDChiNhanh"     json:"IDChiNhanh,omitempty"`
	ChiNhanh       sql.NullString    `db:"ChiNhanh"       json:"ChiNhanh,omitempty"`
	
	IDPhongBan     sql.NullInt64       `db:"IDPhongBan"     json:"IDPhongBan,omitempty"`
	TenPhongBan    sql.NullString       `db:"TenPhongBan"    json:"TenPhongBan,omitempty"`
	
	IDChucDanh     sql.NullInt64       `db:"IDChucDanh"     json:"IDChucDanh,omitempty"`
	TenChucDanh    sql.NullString    `db:"TenChucDanh"    json:"TenChucDanh,omitempty"`
	
	
	NgayBatDau     time.Time `db:"NgayBatDau"     json:"NgayBatDau,omitempty"`
	NgayKetThuc    time.Time `db:"NgayKetThuc"    json:"NgayKetThuc,omitempty"`

	TenNguoiThan    sql.NullString `db:"TenNguoiThan"     json:"TenNguoiThan,omitempty"`
	SDTNguoiThan    sql.NullString `db:"SDTNguoiThan"     json:"SDTNguoiThan,omitempty"`
	DiaChiNguoiThan sql.NullString `db:"DiaChiNguoiThan"  json:"DiaChiNguoiThan,omitempty"`
	QuanHe          sql.NullString `db:"QuanHe"           json:"QuanHe,omitempty"`
}
