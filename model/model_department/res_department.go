package modeldepartment

type ResDepartment struct {
	IDChiNhanh  int    `db:"IDChiNhanh"  json:"IDChiNhanh,omitempty"`
	TenPhongBan string `db:"TenPhongBan"          json:"TenPhongBan,omitempty"`
	ChiNhanh    string `db:"ChiNhanh"          json:"ChiNhanh,omitempty"`
	ID          int    `db:"ID"          json:"ID,omitempty"`
}
