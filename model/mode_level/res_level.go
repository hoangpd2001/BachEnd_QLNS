package modelevel

type ResLevel struct {
	ID           int    `db:"ID"                    json:"ID,omitempty"`
	TenCapBac    string `db:"TenCapBac"             json:"TenCapBac,omitempty"`
	CauTrucLuong int    `db:"CauTrucLuong"          json:"CauTrucLuong,omitempty"`
}
