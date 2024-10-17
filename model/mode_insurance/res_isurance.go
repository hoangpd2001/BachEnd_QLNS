package modeinsurance

import "time"

type ResInsurance struct {
	ID         int    `db:"ID"                 json:"ID,omitempty"`
	TenBaoHiem string `db:"TenBaoHiem"                 json:"TenBaoHiem,omitempty"`
	NhaCungCap string `db:"NhaCungCap"                 json:"NhaCungCap,omitempty"`
	NoiDangKi  string `db:"NoiDangKi"                 json:"NoiDangKi,omitempty"`
	TyLePhi    int    `db:"TyLePhi"                 json:"TyLePhi,omitempty"`
}

type ResUserInsurance struct {
	IDNhanVien int       `db:"IDNhanVien"                 json:"IDNhanVien,omitempty"`
	IDBaoHiem  int       `db:"IDBaoHiem"                 json:"IDBaoHiem,omitempty"`
	NgayDong   time.Time `db:"NgayDong"                 json:"NgayDong,omitempty"`
	NgayHetHan time.Time `db:"NgayHetHan"                 json:"NgayHetHan,omitempty"`
}
