package repobranch

import (
	"BackEnd/mod/banana"
	modelbranch "BackEnd/mod/model/model_Branch"

	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type BranchRepo struct {
	sqlDB *sqlx.DB
}

func NewBranchRepo(sql *sqlx.DB) *BranchRepo {
	return &BranchRepo{
		sqlDB: sql,
	}
}

func (u BranchRepo) CreatBranch(context context.Context, branch modelbranch.ResBranch) (modelbranch.ResBranch, error) {
	statement :=
		`
		INSERT INTO chinhanh(ChiNhanh) VALUES (:ChiNhanh)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, branch)
	if err != nil {
		log.Error(err.Error())
		return branch, banana.UpdateFailed
	}
	return branch, nil
}

func (u BranchRepo) SelectBranchAll(context context.Context) ([]modelbranch.ResBranch, error) {
	var Slicebranch []modelbranch.ResBranch
	query := "SELECT * FROM chinhanh"
	err := u.sqlDB.SelectContext(context, &Slicebranch, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return Slicebranch, nil
}

func (u BranchRepo) SelelectBranchById(context context.Context, BranchId int) (modelbranch.ResBranch, error) {
	var typeUser modelbranch.ResBranch
	query := "SELECT * FROM chinhanh WHERE ID=?"
	err := u.sqlDB.GetContext(context, &typeUser, query, BranchId)
	if err != nil {
		log.Error(err.Error())
		return typeUser, banana.GetIdFailed
	}
	return typeUser, nil
}

func (u BranchRepo) UpdateTypeById(context context.Context, branch modelbranch.ResBranch) (modelbranch.ResBranch, error) {
	statement :=
		`
		UPDATE chinhanh SET ChiNhanh=:ChiNhanh WHERE ID=:ID;
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, branch)
	if err != nil {
		log.Error(err.Error())
		return branch, banana.UpdateFailed
	}
	return branch, nil
}
func (u BranchRepo) DeleteBranchById(context context.Context, BranchId int) (sql.Result, error) {
	query := "DELETE FROM chinhanh WHERE ID = ?"
	result, err := u.sqlDB.ExecContext(context, query, BranchId)
	if err != nil {
		log.Error(err.Error())
		return result, banana.UpdateFailed
	}
	return result, nil
}
