package repotype

import (
	"BackEnd/mod/banana"
	modetypeuser "BackEnd/mod/model/mode_typeUser"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type TypeRepo struct {
	sqlDB *sqlx.DB
}

func NewTypeRepo(sql *sqlx.DB) *TypeRepo {
	return &TypeRepo{
		sqlDB: sql,
	}
}

func (u TypeRepo) CreatTYpe(context context.Context, UserType modetypeuser.ResTypeUser) (modetypeuser.ResTypeUser, error) {
	statement :=
		`
		INSERT INTO loainhanvien(LoaiNhanVien) VALUES (:LoaiNhanVien)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserType)
	if err != nil {
		log.Error(err.Error())
		return UserType, banana.UpdateFailed
	}
	return UserType, nil
}

func (u TypeRepo) SelectTypeAll(context context.Context) ([]modetypeuser.ResTypeUser, error) {
	var SliceType []modetypeuser.ResTypeUser
	query := "SELECT * FROM loainhanvien"
	err := u.sqlDB.SelectContext(context, &SliceType, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceType, nil
}

func (u TypeRepo) SelelectTypeByUser(context context.Context, UserId int) (modetypeuser.ResTypeUser, error) {
	var typeUser modetypeuser.ResTypeUser
	query := "SELECT loainhanvien.ID, loainhanvien.LoaiNhanVien FROM loainhanvien, nhanvien WHERE nhanvien.LoaiNhanVien = loainhanvien.ID and nhanvien.ID=?"
	err := u.sqlDB.GetContext(context, &typeUser, query, UserId)
	if err != nil {
		log.Error(err.Error())
		return typeUser, banana.GetIdFailed
	}
	return typeUser, nil
}

func (u TypeRepo) UpdateTypeById(context context.Context, UserType modetypeuser.ResTypeUser) (modetypeuser.ResTypeUser, error) {
	statement :=
		`
		UPDATE loainhanvien SET LoaiNhanVien=:LoaiNhanVien WHERE ID=:ID
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserType)
	if err != nil {
		log.Error(err.Error())
		return UserType, banana.UpdateFailed
	}
	return UserType, nil
}
func (u TypeRepo) DeleteTypeById(context context.Context, TypeID int) (sql.Result, error) {
	query := "DELETE FROM `loainhanvien` WHERE ID = ?"
	result, err := u.sqlDB.ExecContext(context, query, TypeID)
	if err != nil {
		log.Error(err.Error())
		return result, banana.UpdateFailed
	}
	return result, nil
}
