package modelgrup


type ResGrup struct {
	ID      int    `db:"ID"                 json:"ID,omitempty"`
	TenNhom string `db:"TenNhom"            json:"TenNhom,omitempty"`
}

type ResUserGrup struct {
	IDNhanVien  int       `db:"IDNhanVien"       json:"IDNhanVien,omitempty"`
	IDNhom    int       `db:"IDNhom"         json:"IDNhom,omitempty"`
	Ten            string    `db:"Ten"            json:"Ten,omitempty"`
	Dem            string    `db:"Dem"            json:"Dem,omitempty"`
	Ho             string    `db:"Ho"             json:"Ho,omitempty"`
	TenNhom       string    `db:"TenNhom"            json:"TenNhom,omitempty"`
	
}
