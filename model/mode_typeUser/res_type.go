package modetypeuser

type ResTypeUser struct {
	ID           int    `db:"ID"                    json:"ID,omitempty"`
	LoaiNhanVien string `db:"LoaiNhanVien"          json:"LoaiNhanVien,omitempty"`
}
