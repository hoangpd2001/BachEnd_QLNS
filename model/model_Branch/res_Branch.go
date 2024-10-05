package modelbranch

type ResBranch struct {
	ID       int    `db:"ID"                    json:"ID,omitempty"`
	ChiNhanh string `db:"ChiNhanh"          json:"ChiNhanh,omitempty"`
}
