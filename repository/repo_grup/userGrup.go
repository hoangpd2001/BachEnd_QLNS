package repogrup

import (
	"BackEnd/mod/banana"
	modelgrup "BackEnd/mod/model/model_grup"
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type GrupUserRepo struct {
	sqlDB *sqlx.DB
}

func NewGrupUserRepo(sql *sqlx.DB) *GrupUserRepo {
	return &GrupUserRepo{
		sqlDB: sql,
	}
}

func (u GrupUserRepo) CreatGrupUser(context context.Context, sUserGrup modelgrup.ResUserGrup) (modelgrup.ResUserGrup, error) {
	statement :=
		`
		INSERT INTO nhanvien_nhom(IDNhanVien, IDNhom) VALUES (:IDNhanVien, :IDNhom)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, sUserGrup)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return sUserGrup, banana.SameName
		}
		log.Error(err.Error())
		return sUserGrup, banana.SererError
	}
	return sUserGrup, nil
}

func (u GrupUserRepo) SelectGrupUserByGrup(context context.Context, GrupID int) ([]modelgrup.ResUserGrup, error) {
	var SliceSkill []modelgrup.ResUserGrup
	query := `SELECT  nhanvien_nhom.*,nhanvien.Ten, nhanvien.Dem, nhanvien.Ho, nhomnhanvien.TenNhom FROM nhanvien_nhom, nhomnhanvien, nhanvien WHERE 
	nhanvien_nhom.IDNhanVien = nhanvien.ID
    and nhanvien_nhom.IDNhom = nhomnhanvien.ID
	and nhanvien_nhom.IDNhom = ?`
	err := u.sqlDB.SelectContext(context, &SliceSkill, query, GrupID)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceSkill, nil
}

func (u GrupUserRepo) SelelectGrupUseeAll(context context.Context) ([]modelgrup.ResUserGrup, error) {
	var Skill []modelgrup.ResUserGrup

	query := `SELECT  nhanvien_nhom.*,nhanvien.Ten, nhanvien.Dem, nhanvien.Ho, nhomnhanvien.TenNhom FROM nhanvien_nhom, nhomnhanvien, nhanvien WHERE 
	nhanvien_nhom.IDNhanVien = nhanvien.ID
    and nhanvien_nhom.IDNhom = nhomnhanvien.ID
`
	err := u.sqlDB.SelectContext(context, &Skill, query)
	if err != nil {
		log.Error(err.Error())
		return Skill, banana.GetIdFailed
	}
	return Skill, nil
}

// func (u GrupUserRepo) UpdateSkillById(context context.Context, sUserGrup modelgrup.ResUserGrup) (modelgrup.ResUserGrup, error) {
// 	statement :=
// 		`
// 			UPDATE nhanvien_kynang SET MucDo=:MucDo,
// 			NgayDanhGia=:NgayDanhGia,IDKyNang = :IDKyNangMoi WHERE IDNhanVien = :IDNhanVien and IDKyNang = :IDKyNang		
// 		`
// 	_, err := u.sqlDB.NamedExecContext(context, statement, sUserGrup)
// 	if err != nil {
// 		if err.(*mysql.MySQLError).Number == 1062 {
// 			log.Error(err.Error())
// 			return sUserGrup, banana.SameName
// 		}
// 		log.Error(err.Error())
// 		return sUserGrup, banana.SererError
// 	}
// 	return sUserGrup, nil
// }
func (u GrupUserRepo) DeleteGrupUser(context context.Context, GrupID int, UserId int) (sql.Result, error) {
	query := "DELETE FROM nhanvien_nhom WHERE IDNhanVien = ? and IDNhom = ?"
	result, err := u.sqlDB.ExecContext(context, query, UserId, GrupID)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return result, banana.ForenkeyErrol
		}
		log.Error(err.Error())
		return result, banana.SererError
	}
	return result, nil
}
