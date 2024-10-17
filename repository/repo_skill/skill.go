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

type SkillRepo struct {
	sqlDB *sqlx.DB
}

func NewSkillRepo(sql *sqlx.DB) *SkillRepo {
	return &SkillRepo{
		sqlDB: sql,
	}
}

func (u SkillRepo) CreatSkill(context context.Context, UserSkill modelskill.ResSkill) (modelskill.ResSkill, error) {
	statement :=
		`
		INSERT INTO kynang(TenKyNang, MoTa) VALUES (:TenKyNang, :MoTa)
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

func (u SkillRepo) SelectSkillAll(context context.Context) ([]modelskill.ResSkill, error) {
	var SliceSkill []modelskill.ResSkill
	query := "SELECT * FROM kynang"
	err := u.sqlDB.SelectContext(context, &SliceSkill, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceSkill, nil
}

func (u SkillRepo) SelelectSkillById(context context.Context, SkillId int) (modelskill.ResSkill, error) {
	var Skill modelskill.ResSkill
	query := "SELECT * FROM kynang WHERE ID=?"
	err := u.sqlDB.GetContext(context, &Skill, query, SkillId)
	if err != nil {
		log.Error(err.Error())
		return Skill, banana.GetIdFailed
	}
	return Skill, nil
}

func (u SkillRepo) UpdateSkillById(context context.Context, UserSkill modelskill.ResSkill) (modelskill.ResSkill, error) {
	statement :=
		`
			UPDATE kynang SET TenKyNang=:TenKyNang,MoTa=:MoTa WHERE ID = :ID		
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
func (u SkillRepo) DeleteSkillById(context context.Context, SkillID int) (sql.Result, error) {
	query := "DELETE FROM `kynang` WHERE ID = ?"
	result, err := u.sqlDB.ExecContext(context, query, SkillID)
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
