package modelevel

type ReqLevel struct {
	TenCapBac    string `validate:"required"`
	CauTrucLuong int    `validate:"required"`
}
