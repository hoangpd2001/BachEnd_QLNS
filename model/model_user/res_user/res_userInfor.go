package resuser

type ResUserInfor struct {
	IDNhanVien           int       `db:"IDNhanVien"              json:"-"`
	GioiTinh        bool   `db:"GioiTinh"        json:"GioiTinh,omitempty"`
	SDT             string `db:"SDT"             json:"SDT,omitempty"`
	EmailCaNhan     string `db:"EmailCaNhan"     json:"EmailCaNhan,omitempty"`
	DiaChiThuongTru string `db:"DiaChiThuongTru" json:"DiaChiThuongTru,omitempty"`
	DiaChiTamTru    string `db:"DiaChiTamTru"    json:"DiaChiTamTru,omitempty"`
	CCCD            string `db:"CCCD"            json:"CCCD,omitempty"`
}