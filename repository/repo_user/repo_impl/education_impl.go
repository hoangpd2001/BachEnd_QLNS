package repoimpl

import (
	"BackEnd/mod/banana"
	res_user "BackEnd/mod/model/model_user/res_user"
	repouser "BackEnd/mod/repository/repo_user"
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type EducationRepoImpl struct {
	sqlDB *sqlx.DB
}

// Hàm khởi tạo NewEducationRepo
func NewEducationRepo(sqlDB *sqlx.DB) repouser.EducationRepo {
	return &EducationRepoImpl{
		sqlDB: sqlDB,
	}
}

func (u EducationRepoImpl) CreatEducation(context context.Context, UserEdu res_user.ResEducation) (res_user.ResEducation, error) {
	statement :=
		`
		INSERT INTO nhanvien_hocvan( IDNhanVien, Truong, BangCap, 
		CapHoc, NamTotNghiep) VALUES (:IDNhanVien, :Truong, :BangCap, 
		:CapHoc, :NamTotNghiep)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserEdu)
	if err != nil {
		log.Error(err.Error())
		return UserEdu, banana.UpdateFailed
	}
	return UserEdu, nil
}

func (u EducationRepoImpl) SelectEducationByUser(context context.Context, UserId int) ([]res_user.ResEducation, error) {
	var SliceEducation []res_user.ResEducation
	query := "SELECT * FROM nhanvien_hocvan WHERE IDNhanVien = ?"
	err := u.sqlDB.SelectContext(context, &SliceEducation, query, UserId)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceEducation, nil
}

func (u EducationRepoImpl) SelectEducationById(context context.Context, EducationId int) (res_user.ResEducation, error) {
	var Education res_user.ResEducation
	query := "SELECT * FROM nhanvien_hocvan WHERE ID = ?"
	err := u.sqlDB.GetContext(context, &Education, query, EducationId)
	if err != nil {
		log.Error(err.Error())
		return Education, banana.GetIdFailed
	}
	return Education, nil
}

func (u EducationRepoImpl) UpdateEducationById(context context.Context, UserEdu res_user.ResEducation) (res_user.ResEducation, error) {
	statement :=
		`
			UPDATE nhanvien_hocvan SET 
				Truong=:Truong,BangCap=:BangCap,CapHoc=:CapHoc,NamTotNghiep=:NamTotNghiep
		 	WHERE ID= :ID
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserEdu)
	if err != nil {
		log.Error(err.Error())
		return UserEdu, banana.UpdateFailed
	}
	return UserEdu, nil
}

func (u EducationRepoImpl) DeleteEducationById(context context.Context, UserId int) (sql.Result, error) {

	query := "Delete from nhanvien_hocvan where ID = ?"
	result, err := u.sqlDB.ExecContext(context, query, UserId)
	if err != nil {
		log.Error(err.Error())
		return result, banana.UpdateFailed
	}
	return result, nil
}
