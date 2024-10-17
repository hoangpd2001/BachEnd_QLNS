package repotitle

import (
	"BackEnd/mod/banana"
	modeltitle "BackEnd/mod/model/model_title"
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type TitleRepo struct {
	sqlDB *sqlx.DB
}

func NewTitleRepo(sql *sqlx.DB) *TitleRepo {
	return &TitleRepo{
		sqlDB: sql,
	}
}

func (u TitleRepo) CreatTitle(context context.Context, Title modeltitle.ResTitle) (modeltitle.ResTitle, error) {
	statement :=
		`
		INSERT INTO chucdanh(TenChucDanh) VALUES (:TenChucDanh)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, Title)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return Title, banana.SameName
		}
		log.Error(err.Error())
		return Title, banana.SererError
	}
	return Title, nil
}

func (u TitleRepo) SelectTitleAll(context context.Context) ([]modeltitle.ResTitle, error) {
	var SliceTitle []modeltitle.ResTitle
	query := "SELECT * FROM chucdanh"
	err := u.sqlDB.SelectContext(context, &SliceTitle, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceTitle, nil
}

func (u TitleRepo) SelelectTitleById(context context.Context, TitleId int) (modeltitle.ResTitle, error) {
	var typeUser modeltitle.ResTitle
	query := "SELECT * FROM chucdanh WHERE ID=?"
	err := u.sqlDB.GetContext(context, &typeUser, query, TitleId)
	if err != nil {
		log.Error(err.Error())
		return typeUser, banana.GetIdFailed
	}
	return typeUser, nil
}

func (u TitleRepo) UpdateTitleById(context context.Context, Title modeltitle.ResTitle) (modeltitle.ResTitle, error) {
	statement :=
		`
		UPDATE chucdanh SET TenChucDanh=:TenChucDanh WHERE ID=:ID;
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, Title)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062{
		log.Error(err.Error())
		return Title, banana.SameName
		}
		log.Error(err.Error())
		return Title, banana.SererError
	}
	return Title, nil
}
func (u TitleRepo) DeleteTitleById(context context.Context, TitleId int) (sql.Result, error) {
	query := "DELETE FROM chucdanh WHERE ID = ?"
	result, err := u.sqlDB.ExecContext(context, query, TitleId)
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
