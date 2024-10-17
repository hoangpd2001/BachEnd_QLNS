package repoinsurance


import (
	"BackEnd/mod/banana"
	modelinsurance "BackEnd/mod/model/mode_insurance"
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type InsuranceUserRepo struct {
	sqlDB *sqlx.DB
}

func NewInsuranceUserRepo(sql *sqlx.DB) *InsuranceUserRepo {
	return &InsuranceUserRepo{
		sqlDB: sql,
	}
}

func (u InsuranceUserRepo) CreatInsuranceUser(context context.Context, Insurance modelinsurance.ResUserInsurance) (modelinsurance.ResUserInsurance, error) {
	statement :=
		`
		INSERT INTO nhanvien_baohiem(IDNhanVien, IDBaoHiem, NgayDong, NgayHetHan) 
		VALUES (:IDNhanVien, :IDBaoHiem, :NgayDong, :NgayHetHan)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, Insurance)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return Insurance, banana.SameName
		}
		log.Error(err.Error())
		return Insurance, banana.SererError
	}
	return Insurance, nil
}

func (u InsuranceUserRepo) SelectInsuranceUserAll(context context.Context) ([]modelinsurance.ResUserInsurance, error) {
	var SliceInsurance []modelinsurance.ResUserInsurance
	query := "SELECT * FROM `nhanvien_baohiem`"
	err := u.sqlDB.SelectContext(context, &SliceInsurance, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceInsurance, nil
}
func (u InsuranceUserRepo) SelectInsuranceUserByUser(context context.Context, IDUser int) ([]modelinsurance.ResUserInsurance, error) {
	var SliceInsurance []modelinsurance.ResUserInsurance
	query := " SELECT * FROM `nhanvien_baohiem` WHERE IDNhanVien = ?"
	err := u.sqlDB.SelectContext(context, &SliceInsurance, query,IDUser)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceInsurance, nil
}

func (u InsuranceUserRepo) SelelectInsuranceUserByOne(context context.Context, InsuranceId int, UserId int) (modelinsurance.ResUserInsurance, error) {
	var Insurance modelinsurance.ResUserInsurance

	query := `SELECT * FROM nhanvien_baohiem WHERE IDNhanVien = ? and IDBaoHiem = ?`
	err := u.sqlDB.GetContext(context, &Insurance, query, UserId, InsuranceId)
	if err != nil {
		
		log.Error(err.Error())
		return Insurance, banana.GetIdFailed
	}
	return Insurance, nil
}

func (u InsuranceUserRepo) UpdateInsuranceById(context context.Context, Insurance modelinsurance.ResUserInsurance) (modelinsurance.ResUserInsurance, error) {
	statement :=
		`
			UPDATE nhanvien_baohiem SET 
			NgayDong=:NgayDong ,NgayHetHan=:NgayHetHan 
			 WHERE 
			 IDNhanVien=:IDNhanVien and  IDBaoHiem=:IDBaoHiem 	
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, Insurance)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return Insurance, banana.SameName
		}
		log.Error(err.Error())
		return Insurance, banana.SererError
	}
	return Insurance, nil
}
func (u InsuranceUserRepo) DeleteInsuranceById(context context.Context, IDBH int, UserId int) (sql.Result, error) {
	query := "DELETE FROM nhanvien_baohiem WHERE IDNhanVien = ? and IDBaoHiem = ?"
	result, err := u.sqlDB.ExecContext(context, query, UserId, IDBH)
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
