package resuser

type ResUserCount struct {
	Thang             string       `db:"Thang"             json:"Thang,omitempty"`
	SoLuong            int    `db:"SoLuong"            json:"SoLuong"`
}