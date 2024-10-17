package repotitle

import (
	"BackEnd/mod/banana"
	modeltitle "BackEnd/mod/model/model_title"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type UserTitleRepo struct {
	sqlDB *sqlx.DB
}

func NewUserTitleRepo(sql *sqlx.DB) *UserTitleRepo {
	return &UserTitleRepo{
		sqlDB: sql,
	}
}

func (u UserTitleRepo) CreatUserTitle(context context.Context, UserTitle modeltitle.ResUserTitle) (modeltitle.ResUserTitle, error) {
	statement :=
		`
		INSERT INTO nhanvien_chucdanh(IDNhanVien, IDChucDanh, NgayBatDau, NgayKetThuc, IDPhongBan)
		 VALUES (:IDNhanVien, :IDChucDanh, :NgayBatDau, :NgayKetThuc, :IDPhongBan)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserTitle)
	if err != nil {
		log.Error(err.Error())
		return UserTitle, banana.UpdateFailed
	}
	return UserTitle, nil
}

func (u UserTitleRepo) SelectUserTitleAll(context context.Context, IdUser int) ([]modeltitle.ResUserTitle, error) {
	var SliceUserTitle []modeltitle.ResUserTitle
	query := `SELECT IDNhanVien, IDChucDanh, IDChiNhanh,nhanvien_chucdanh.NgayBatDau, nhanvien_chucdanh.NgayKetThuc,
	 IDPhongBan, nhanvien.Ho,nhanvien.Dem,nhanvien.Ten, chucdanh.TenChucDanh ,
	 phongban.TenPhongBan FROM nhanvien_chucdanh, nhanvien, chucdanh, phongban,chinhanh
	  WHERE nhanvien_chucdanh.IDNhanVien = nhanvien.ID and nhanvien_chucdanh.IDChucDanh=chucdanh.ID
	  and chinhanh.ID = phongban.IDChiNhanh 
	  and nhanvien_chucdanh.IDPhongBan =phongban.ID and IDNhanVien = ?`
	err := u.sqlDB.SelectContext(context, &SliceUserTitle, query,IdUser)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceUserTitle, nil
}

func (u UserTitleRepo) SelelectTitleByTitle(context context.Context, TitleId int,IdUser int ) (modeltitle.ResUserTitle, error) {
	var userTitle modeltitle.ResUserTitle
	query := `SELECT IDNhanVien, IDChucDanh, nhanvien_chucdanh.NgayBatDau, nhanvien_chucdanh.NgayKetThuc,
	 IDPhongBan, nhanvien.Ho,nhanvien.Dem,nhanvien.Ten, chucdanh.TenChucDanh ,
	 phongban.TenPhongBan FROM nhanvien_chucdanh, nhanvien, chucdanh, phongban
	  WHERE nhanvien_chucdanh.IDNhanVien = nhanvien.ID and nhanvien_chucdanh.IDChucDanh=chucdanh.ID 
	  and nhanvien_chucdanh.IDPhongBan =phongban.ID and IDNhanVien = ? and IDChucDanh = ?`
	err := u.sqlDB.GetContext(context, &userTitle, query,IdUser, TitleId)
	if err != nil {
		log.Error(err.Error())
		return userTitle, banana.GetIdFailed
	}
	return userTitle, nil
}

func (u UserTitleRepo) UpdateTitleById(context context.Context, Title modeltitle.ResUserTitle) (modeltitle.ResUserTitle, error) {
	statement :=
		`
		UPDATE nhanvien_chucdanh SET NgayBatDau=:NgayBatDau,NgayKetThuc=:NgayKetThuc WHERE IDNhanVien=:IDNhanVien and IDChucDanh=:IDChucDanh and IDPhongBan=:IDPhongBan
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, Title)
	if err != nil {
		log.Error(err.Error())
		return Title, banana.UpdateFailed
	}
	return Title, nil
}
func (u UserTitleRepo) DeleteTitleById(context context.Context, TitleId int,IdUser int, idpb int) (sql.Result, error) {
	query := "DELETE FROM nhanvien_chucdanh WHERE IDNhanVien=? and IDChucDanh=? and IDPhongBan=?"  
	result, err := u.sqlDB.ExecContext(context, query,IdUser, TitleId, idpb)
	if err != nil {
		log.Error(err.Error())
		return result, banana.UpdateFailed
	}
	return result, nil
}
