package repoimpl

import (
	"BackEnd/mod/banana"
	res_user "BackEnd/mod/model/model_user/res_user"
	repouser "BackEnd/mod/repository/repo_user"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type RelativeRepoImpl struct {
	sqlDB *sqlx.DB
}

func NewRelativeRepo(sql *sqlx.DB) repouser.RelativeRepo {
	return &RelativeRepoImpl{
		sqlDB: sql,
	}
}

func (u RelativeRepoImpl) CreatRelative(context context.Context, UserRelative res_user.ResRelative) (res_user.ResRelative, error) {
	statement :=
		`
		INSERT INTO nhanvien_nguoithan(IDNhanVien, TenNguoiThan, SDTNguoiThan, QuanHe, DiaChiNguoiThan)
		 VALUES (:IDNhanVien,:TenNguoiThan,:SDTNguoiThan,:QuanHe,:DiaChiNguoiThan)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserRelative)
	if err != nil {
		log.Error(err.Error())
		return UserRelative, banana.UpdateFailed
	}
	return UserRelative, nil
}

func (u RelativeRepoImpl) SelectRelativeByUser(context context.Context, UserId int) ([]res_user.ResRelative, error) {
	var SliceRelative []res_user.ResRelative
	query := "SELECT * FROM nhanvien_nguoithan WHERE IDNhanVien = ?"
	err := u.sqlDB.SelectContext(context, &SliceRelative, query, UserId)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceRelative, nil
}

func (u RelativeRepoImpl) UpdateRelativeByUser(context context.Context, UserRelative res_user.ResRelative) (res_user.ResRelative, error) {
	statement :=
		`
		UPDATE nhanvien_nguoithan SET TenNguoiThan=:TenNguoiThan,
		SDTNguoiThan=:SDTNguoiThan,QuanHe=:QuanHe,DiaChiNguoiThan=:DiaChiNguoiThan WHERE IDNhanVien=:IDNhanVien 
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserRelative)
	if err != nil {
		log.Error(err.Error())
		return UserRelative, banana.UpdateFailed
	}
	return UserRelative, nil
}
