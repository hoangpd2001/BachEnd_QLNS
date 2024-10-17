package repodepartment

import (
	"BackEnd/mod/banana"
	modeldepartment "BackEnd/mod/model/model_department"
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type Department struct {
	sqlDB *sqlx.DB
}

func NewDepartment(sql *sqlx.DB) *Department {
	return &Department{
		sqlDB: sql,
	}
}

func (u Department) CreatDepartment(context context.Context, Department modeldepartment.ResDepartment) (modeldepartment.ResDepartment, error) {
	statement :=
		`
		INSERT INTO phongban( TenPhongBan, IDChiNhanh) VALUES (:TenPhongBan, :IDChiNhanh)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, Department)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062{
		log.Error(err.Error())
		return Department, banana.SameName
	}
		log.Error(err.Error())
		return Department, banana.UpdateFailed
	}
	return Department, nil
}

func (u Department) SelectDepartmentAll(context context.Context) ([]modeldepartment.ResDepartment, error) {
	var SliceDepartment []modeldepartment.ResDepartment
	query := "SELECT phongban.*, ChiNhanh FROM phongban,chinhanh  WHERE phongban.IDChiNhanh = chinhanh.ID "
	err := u.sqlDB.SelectContext(context, &SliceDepartment, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceDepartment, nil
}
func (u Department) SelectDepartmentByBranch(context context.Context, idBranch int) ([]modeldepartment.ResDepartment, error) {
	var SliceDepartment []modeldepartment.ResDepartment
	query := "SELECT * FROM phongban  WHERE IDChiNhanh = ? "
	err := u.sqlDB.SelectContext(context, &SliceDepartment, query, idBranch)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceDepartment, nil
}
func (u Department) SelelectDepartmentById(context context.Context, DepartmentID int) (modeldepartment.ResDepartment, error) {
	var Department modeldepartment.ResDepartment
	
	query := `SELECT phongban.*, ChiNhanh FROM phongban,chinhanh  WHERE phongban.IDChiNhanh = chinhanh.ID and phongban.ID=?`
	err := u.sqlDB.GetContext(context, &Department, query, DepartmentID)
	if err != nil {
		log.Error(err.Error())
		return Department, banana.GetIdFailed
	}
	return Department, nil
}

func (u Department) UpdateSkillById(context context.Context, Department modeldepartment.ResDepartment) (modeldepartment.ResDepartment, error) {
 	statement :=
		`
			UPDATE phongban SET TenPhongBan=:TenPhongBan,IDChiNhanh=:IDChiNhanh WHERE ID=:ID		
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, Department)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062{
		log.Error(err.Error())
		return Department, banana.SameName
	}
		log.Error(err.Error())
		return Department, banana.SererError
	}
	return Department, nil
}
func (u Department) DeleteSkillById(context context.Context, DepartmentID int) (sql.Result, error) {
	query := "DELETE FROM phongban WHERE ID=?"
	result, err := u.sqlDB.ExecContext(context, query, DepartmentID)
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
