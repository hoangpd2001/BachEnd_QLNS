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
	NgayBatDau   time.Time `db:"NgayBatDau"     json:"NgayBatDau,omitempty"`
	NgayKetThuc  time.Time `db:"NgayKetThuc"    json:"NgayKetThuc,omitempty"`
	Token        string    `                    json:"token,omitempty"`
}
