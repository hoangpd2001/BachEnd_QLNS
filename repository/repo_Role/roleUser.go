package reporole

import (
	"BackEnd/mod/banana"
	modeRole "BackEnd/mod/model/model_role"
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type RoleUserRepo struct {
	sqlDB *sqlx.DB
}

func NewRoleUserRepo(sql *sqlx.DB) *RoleUserRepo {
	return &RoleUserRepo{
		sqlDB: sql,
	}
}

func (u RoleUserRepo) CreatRoleUser(context context.Context, UserRole modeRole.ResUserRole) (modeRole.ResUserRole, error) {
	statement :=
		`
		INSERT INTO chucdanh_vaitro(IDChucDanh, IDVaiTro, Xem, Them, Sua, Xoa)
		 VALUES (:IDChucDanh, :IDVaiTro, :Xem, :Them, :Sua, :Xoa)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserRole)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return UserRole, banana.SameName
		}
		log.Error(err.Error())
		return UserRole, banana.SererError
	}
	return UserRole, nil
}

func (u RoleUserRepo) SelectRoleUserAll(context context.Context) ([]modeRole.ResUserRole, error) {
	var SliceSkill []modeRole.ResUserRole
	query := "SELECT * FROM chucdanh_vaitro"
	err := u.sqlDB.SelectContext(context, &SliceSkill, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceSkill, nil
}
// func (u RoleUserRepo) SelectRoleUser(context context.Context) ([]modeRole.ResUserRole, error) {
// 	var SliceSkill []modeRole.ResUserRole
// 	query := "SELECT nhanvien_kynang.*,nhanvien.Ten, nhanvien.Dem, nhanvien.Ho, TenKyNang FROM nhanvien_kynang, nhanvien,kynang WHERE nhanvien.ID = nhanvien_kynang.IDNhanVien and kynang.ID = nhanvien_kynang.IDKyNang "
// 	err := u.sqlDB.SelectContext(context, &SliceSkill, query)
// 	if err != nil {
// 		log.Error(err.Error())
// 		return nil, banana.GetIdFailed
// 	}
// 	return SliceSkill, nil
// }

func (u RoleUserRepo) SelelectRoleUser(context context.Context, TitleID []int,RoleID int) (modeRole.ResUserRole, error) {
	var Skill modeRole.ResUserRole

	query := `SELECT 
					MAX(Xem) AS Xem,
					MAX(Them) AS Them,
					MAX(Sua) AS Sua,
					MAX(Xoa) AS Xoa
				FROM 
					chucdanh_vaitro
				WHERE 
					IDVaiTro = ?
					AND IDChucDanh IN (?)
	`
	query, args, err := sqlx.In(query, RoleID, TitleID)
	if err != nil {
		log.Error(err.Error())
		return Skill, banana.GetIdFailed
	}
	query = u.sqlDB.Rebind(query)
	err = u.sqlDB.GetContext(context, &Skill, query, args...)
	if err != nil {
		log.Error(err.Error())
		return Skill, banana.GetIdFailed
	}

	return Skill, nil
}

func (u RoleUserRepo) UpdateSkillById(context context.Context, UserRole modeRole.ResUserRole) (modeRole.ResUserRole, error) {
	statement :=
		`
			UPDATE chucdanh_vaitro SET Xem=:Xem,
			Them=:Them,Sua=:Sua,Xoa=:Xoa WHERE IDChucDanh=:IDChucDanh and IDVaiTro=:IDVaiTro
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserRole)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return UserRole, banana.SameName
		}
		log.Error(err.Error())
		return UserRole, banana.SererError
	}
	return UserRole, nil
}
func (u RoleUserRepo) DeleteSkillById(context context.Context, IDTitle int, IDRole int) (sql.Result, error) {
	query := "DELETE FROM chucdanh_vaitro WHERE  IDChucDanh=? and IDVaiTro=?"
	result, err := u.sqlDB.ExecContext(context, query, IDTitle , IDRole)
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
