package reposkill

import (
	"BackEnd/mod/banana"
	modelskill "BackEnd/mod/model/model_skill"
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type SkillUserRepo struct {
	sqlDB *sqlx.DB
}

func NewSkillUserRepo(sql *sqlx.DB) *SkillUserRepo {
	return &SkillUserRepo{
		sqlDB: sql,
	}
}

func (u SkillUserRepo) CreatSkillUser(context context.Context, UserSkill modelskill.ResUserSkill) (modelskill.ResUserSkill, error) {
	statement :=
		`
		INSERT INTO nhanvien_kynang(IDNhanVien, IDKyNang, MucDo, NgayDanhGia)
		 VALUES (:IDNhanVien, :IDKyNang, :MucDo, :NgayDanhGia)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserSkill)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return UserSkill, banana.SameName
		}
		log.Error(err.Error())
		return UserSkill, banana.SererError
	}
	return UserSkill, nil
}

func (u SkillUserRepo) SelectSkillUserAll(context context.Context, UserId int) ([]modelskill.ResUserSkill, error) {
	var SliceSkill []modelskill.ResUserSkill
	query := "SELECT nhanvien_kynang.*,nhanvien.Ten, nhanvien.Dem, nhanvien.Ho, TenKyNang FROM nhanvien_kynang, nhanvien,KyNang WHERE nhanvien.ID = nhanvien_kynang.IDNhanVien and kynang.ID = nhanvien_kynang.IDKyNang and nhanvien.ID = ?"
	err := u.sqlDB.SelectContext(context, &SliceSkill, query, UserId)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceSkill, nil
}
func (u SkillUserRepo) SelectSkillUser(context context.Context) ([]modelskill.ResUserSkill, error) {
	var SliceSkill []modelskill.ResUserSkill
	query := "SELECT nhanvien_kynang.*,nhanvien.Ten, nhanvien.Dem, nhanvien.Ho, TenKyNang FROM nhanvien_kynang, nhanvien,kynang WHERE nhanvien.ID = nhanvien_kynang.IDNhanVien and kynang.ID = nhanvien_kynang.IDKyNang "
	err := u.sqlDB.SelectContext(context, &SliceSkill, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceSkill, nil
}

func (u SkillUserRepo) SelelectSkillUser(context context.Context, SkillId int, UserId int) (modelskill.ResUserSkill, error) {
	var Skill modelskill.ResUserSkill

	query := `SELECT nhanvien_kynang.*,nhanvien.Ten, nhanvien.Dem, nhanvien.Ho, TenKyNang, MoTa
	 FROM nhanvien_kynang, nhanvien,kynang WHERE 
	nhanvien.ID = nhanvien_kynang.IDNhanVien and 
	kynang.ID= nhanvien_kynang.IDKyNang and nhanvien.ID = ? and kynang.ID = ?`
	err := u.sqlDB.GetContext(context, &Skill, query, UserId, SkillId)
	if err != nil {
		
		log.Error(err.Error())
		return Skill, banana.GetIdFailed
	}
	return Skill, nil
}

func (u SkillUserRepo) UpdateSkillById(context context.Context, UserSkill modelskill.ResUserSkill) (modelskill.ResUserSkill, error) {
	statement :=
		`
			UPDATE nhanvien_kynang SET MucDo=:MucDo,
			NgayDanhGia=:NgayDanhGia WHERE IDNhanVien = :IDNhanVien and IDKyNang = :IDKyNang		
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserSkill)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return UserSkill, banana.SameName
		}
		log.Error(err.Error())
		return UserSkill, banana.SererError
	}
	return UserSkill, nil
}
func (u SkillUserRepo) DeleteSkillById(context context.Context, SkillID int, UserId int) (sql.Result, error) {
	query := "DELETE FROM nhanvien_kynang WHERE IDNhanVien = ? and IDKyNang = ?"
	result, err := u.sqlDB.ExecContext(context, query, UserId, SkillID)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1451{
		log.Error(err.Error())
		return result, banana.ForenkeyErrol
		}
		log.Error(err.Error())
		return result, banana.UpdateFailed
	}
	return result, nil
}
