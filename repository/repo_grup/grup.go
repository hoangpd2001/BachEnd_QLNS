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

type GrupRepo struct {
	sqlDB *sqlx.DB
}

func NewGrupRepo(sql *sqlx.DB) *GrupRepo {
	return &GrupRepo{
		sqlDB: sql,
	}
}

func (u GrupRepo) CreatGrup(context context.Context, NewGrupRepo modelgrup.ResGrup) (modelgrup.ResGrup, error) {
	statement :=
		`
		INSERT INTO nhomnhanvien(TenNhom) VALUES (:TenNhom)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, NewGrupRepo)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.Error(err.Error())
			return NewGrupRepo, banana.SameName
		}
		log.Error(err.Error())
		return NewGrupRepo, banana.SererError
	}
	return NewGrupRepo, nil
}

func (u GrupRepo) SelectGrupAll(context context.Context) ([]modelgrup.ResGrup, error) {
	var SliceGrup []modelgrup.ResGrup
	query := "SELECT * FROM nhomnhanvien"
	err := u.sqlDB.SelectContext(context, &SliceGrup, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceGrup, nil
}

func (u GrupRepo) UpdateGrup(context context.Context, Grup modelgrup.ResGrup) (modelgrup.ResGrup, error) {
	statement :=
		`
		UPDATE nhomnhanvien SET TenNhom=:TenNhom WHERE ID=:ID;
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, Grup)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062{
		log.Error(err.Error())
		return Grup, banana.SameName
		}
		log.Error(err.Error())
		return Grup, banana.SererError
	}
	return Grup, nil
}
func (u GrupRepo) DeleteGrup(context context.Context, GrupId int) (sql.Result, error) {
	query := "DELETE FROM nhomnhanvien WHERE ID = ?"
	result, err := u.sqlDB.ExecContext(context, query, GrupId)
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
