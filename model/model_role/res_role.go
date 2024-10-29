package modelrole

type ResUserRole struct {
	IDChucDanh int `db:"IDChucDanh"                 json:"IDChucDanh,omitempty"`
	IDVaiTro   int `db:"IDVaiTro"                 json:"IDVaiTro,omitempty"`
	Xem        bool   `db:"Xem"                 json:"Xem"`
	Them       bool   `db:"Them"                 json:"Them"`
	Sua        bool   `db:"Sua"                 json:"Sua"`
	Xoa        bool   `db:"Xoa"                 json:"Xoa"`
}

type ResRole struct {
	ID  int    `db:"ID"                 json:"ID,omitempty"`
	Ten string `db:"Ten"                 json:"Ten,omitempty"`
}
