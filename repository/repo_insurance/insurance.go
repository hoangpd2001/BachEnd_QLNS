package repoinsurance

import (
	"BackEnd/mod/banana"
	modelInsurance "BackEnd/mod/model/mode_insurance"
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type InsuranceRepo struct {
	sqlDB *sqlx.DB
}

func NewInsuranceRepo(sql *sqlx.DB) *InsuranceRepo {
	return &InsuranceRepo{
		sqlDB: sql,
	}
}

func (u InsuranceRepo) CreatInsurance(context context.Context, UserInsurance modelInsurance.ResInsurance) (modelInsurance.ResInsurance, error) {
	statement :=
		`
		INSERT INTO baohiem( TenBaoHiem, NhaCungCap, NoiDangKi, TyLePhi)
		 VALUES ( :TenBaoHiem, :NhaCungCap, :NoiDangKi, :TyLePhi)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserInsurance)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return UserInsurance, banana.SameName
		}
		log.Error(err.Error())
		return UserInsurance, banana.SererError
	}
	return UserInsurance, nil
}

func (u InsuranceRepo) SelectInsuranceAll(context context.Context) ([]modelInsurance.ResInsurance, error) {
	var SliceInsurance []modelInsurance.ResInsurance
	query := "SELECT * FROM baohiem"
	err := u.sqlDB.SelectContext(context, &SliceInsurance, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceInsurance, nil
}

func (u InsuranceRepo) SelelectInsuranceById(context context.Context, InsuranceId int) (modelInsurance.ResInsurance, error) {
	var Insurance modelInsurance.ResInsurance
	query := "SELECT * FROM baohiem WHERE ID=?"
	err := u.sqlDB.GetContext(context, &Insurance, query, InsuranceId)
	if err != nil {
		log.Error(err.Error())
		return Insurance, banana.GetIdFailed
	}
	return Insurance, nil
}

func (u InsuranceRepo) UpdateInsuranceById(context context.Context, UserInsurance modelInsurance.ResInsurance) (modelInsurance.ResInsurance, error) {
	statement :=
		`
			UPDATE baohiem SET
			TenBaoHiem=:TenBaoHiem,NhaCungCap=:NhaCungCap,NoiDangKi=:NoiDangKi,TyLePhi=:TyLePhi
			WHERE ID= :ID	`
	_, err := u.sqlDB.NamedExecContext(context, statement, UserInsurance)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return UserInsurance, banana.SameName
		}
		log.Error(err.Error())
		return UserInsurance, banana.SererError
	}
	return UserInsurance, nil
}
func (u InsuranceRepo) DeleteInsuranceById(context context.Context, InsuranceID int) (sql.Result, error) {
	query := "DELETE FROM `baohiem` WHERE ID = ?"
	result, err := u.sqlDB.ExecContext(context, query, InsuranceID)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1451{
		log.Error(err.Error())
		return result, banana.ForenkeyErrol
		}
		log.Error(err.Error())
		return result, banana.SererError
	}
	return result, nil
}
