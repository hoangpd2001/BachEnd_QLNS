package resuser

type ResEducation struct {
	ID           int    `db:"ID"                json:"ID,omitempty"`
	IDNhanVien   int    `db:"IDNhanVien"        json:"IDNhanVien,omitempty"`
	Truong       string `db:"Truong"            json:"Truong,omitempty"`
	BangCap      string `db:"BangCap"           json:"BangCap,omitempty"`
	CapHoc       string `db:"CapHoc"            json:"CapHoc,omitempty"`
	NamTotNghiep string `db:"NamTotNghiep"      json:"NamTotNghiep,omitempty"`
}
