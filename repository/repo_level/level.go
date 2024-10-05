package repolevel

import (
	"BackEnd/mod/banana"
	modelevel "BackEnd/mod/model/mode_level"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type LevelRepo struct {
	sqlDB *sqlx.DB
}

func NewLevelRepo(sql *sqlx.DB) *LevelRepo {
	return &LevelRepo{
		sqlDB: sql,
	}
}

func (u LevelRepo) CreatLevel(context context.Context, level modelevel.ResLevel) (modelevel.ResLevel, error) {
	statement :=
		`
		INSERT INTO capbac(TenCapBac,CauTrucLuong) VALUES (:TenCapBac,:CauTrucLuong)
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, level)
	if err != nil {
		log.Error(err.Error())
		return level, banana.UpdateFailed
	}
	return level, nil
}

func (u LevelRepo) SelectLevelAll(context context.Context) ([]modelevel.ResLevel, error) {
	var SliceLevel []modelevel.ResLevel
	query := "SELECT * FROM capbac"
	err := u.sqlDB.SelectContext(context, &SliceLevel, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceLevel, nil
}

func (u LevelRepo) SelelectLevelByUser(context context.Context, UserId int) (modelevel.ResLevel, error) {
	var level modelevel.ResLevel
	query := "SELECT * FROM capbac, nhanvien WHERE nhanvien.CapBac = capbac.ID and nhanvien.ID=?"
	err := u.sqlDB.GetContext(context, &level, query, UserId)
	if err != nil {
		log.Error(err.Error())
		return level, banana.GetIdFailed
	}
	return level, nil
}

func (u LevelRepo) UpdateLevelById(context context.Context, level modelevel.ResLevel) (modelevel.ResLevel, error) {
	statement :=
		`
		UPDATE capbac SET TenCapBac=:TenCapBac,CauTrucLuong=:CauTrucLuong WHERE ID=:ID
		`
	_, err := u.sqlDB.NamedExecContext(context, statement, level)
	if err != nil {
		log.Error(err.Error())
		return level, banana.UpdateFailed
	}
	return level, nil
}
func (u LevelRepo) DeleteLevelById(context context.Context, LevelID int) (sql.Result, error) {
	query := "DELETE FROM capbac WHERE ID = ?"
	result, err := u.sqlDB.ExecContext(context, query, LevelID)
	if err != nil {
		log.Error(err.Error())
		return result, banana.UpdateFailed
	}
	return result, nil
}
