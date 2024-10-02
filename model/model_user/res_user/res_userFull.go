package resuser

import "time"

type ResUserFull struct {
	ID           int       `db:"ID"              json:"-"`
	Ten          string    `db:"Ten"             json:"Ten,omitempty"`
	Dem          string    `db:"Dem"             json:"Dem,omitempty"`
	Ho           string    `db:"Ho"              json:"Ho,omitempty"`
	Email        string    `db:"Email"           json:"Email,omitempty"`
	LoaiNhanVien int       `db:"LoaiNhanVien"    json:"LoaiNhanVien,omitempty"`
	CapBac       int       `db:"CapBac"          json:"CapBac,omitempty"`
	ChiNhanh     int       `db:"ChiNhanh"        json:"ChiNhanh,omitempty"`
	NgayBatDau   time.Time `db:"NgayBatDau"      json:"NgayBatDau,omitempty"`
	NgayKetThuc  time.Time `db:"NgayKetThuc"     json:"NgayKetThuc,omitempty"`

	//thongtin
	IDNhanVien      int    `db:"IDNhanVien"      json:"-"`
	GioiTinh        bool   `db:"GioiTinh"        json:"GioiTinh,omitempty"`
	SDT             string `db:"SDT"             json:"SDT,omitempty"`
	EmailCaNhan     string `db:"EmailCaNhan"     json:"EmailCaNhan,omitempty"`
	DiaChiThuongTru string `db:"DiaChiThuongTru" json:"DiaChiThuongTru,omitempty"`
	DiaChiTamTru    string `db:"DiaChiTamTru"    json:"DiaChiTamTru,omitempty"`
	CCCD            string `db:"CCCD"            json:"CCCD,omitempty"`

	//NguoiThan
	TenNguoiThan    string `db:"TenNguoiThan"    json:"TenNguoiThan,omitempty"`
	SDTNguoiThan    string `db:"SDTNguoiThan"    json:"SDTNguoiThan,omitempty"`
	DiaChiNguoiThan string `db:"DiaChiNguoiThan" json:"DiaChiNguoiThan,omitempty"`
	QuanHe          string `db:"QuanHe"          json:"QuanHe,omitempty"`

	Token string `                    json:"token,omitempty"`
}
