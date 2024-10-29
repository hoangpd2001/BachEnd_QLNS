package reporole

import (
	"BackEnd/mod/banana"
	modelRole "BackEnd/mod/model/model_role"
	"context"
	// "database/sql"

	// "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type RoleRepo struct {
	sqlDB *sqlx.DB
}

func NewRole(sql *sqlx.DB) *RoleRepo {
	return &RoleRepo{
		sqlDB: sql,
	}
}

// func (u Role) CreatRole(context context.Context, Role modelRole.ReqUserRole) (modelRole.ReqUserRole, error) {
// 	statement :=
// 		`
// 		INSERT INTO kynang(TenKyNang, MoTa) VALUES (:TenKyNang, :MoTa)
// 	`
// 	_, err := u.sqlDB.NamedExecContext(context, statement, Role)
// 	if err != nil {
// 		if err.(*mysql.MySQLError).Number == 1062 {
// 			log.Error(err.Error())
// 			return Role, banana.SameName
// 		}
// 		log.Error(err.Error())
// 		return Role, banana.SererError
// 	}
// 	return Role, nil
// }

func (u RoleRepo) SelectRoleAll(context context.Context) ([]modelRole.ResRole, error) {
	var SliceRole []modelRole.ResRole
	query := "SELECT * FROM vaitro"
	err := u.sqlDB.SelectContext(context, &SliceRole, query)
	if err != nil {
		log.Error(err.Error())
		return nil, banana.GetIdFailed
	}
	return SliceRole, nil
}

// func (u Role) SelelectRoleById(context context.Context, RoleId int) (modelRole.ReqUserRole, error) {
// 	var Role modelRole.ReqUserRole
// 	query := "SELECT * FROM kynang WHERE ID=?"
// 	err := u.sqlDB.GetContext(context, &Role, query, RoleId)
// 	if err != nil {
// 		log.Error(err.Error())
// 		return Role, banana.GetIdFailed
// 	}
// 	return Role, nil
// }

// func (u Role) UpdateRoleById(context context.Context, Role modelRole.ReqUserRole) (modelRole.ReqUserRole, error) {
// 	statement :=
// 		`
// 			UPDATE kynang SET TenKyNang=:TenKyNang,MoTa=:MoTa WHERE ID = :ID		
// 		`
// 	_, err := u.sqlDB.NamedExecContext(context, statement, Role)
// 	if err != nil {
// 		if err.(*mysql.MySQLError).Number == 1062 {
// 			log.Error(err.Error())
// 			return Role, banana.SameName
// 		}
// 		log.Error(err.Error())
// 		return Role, banana.SererError
// 	}
// 	return Role, nil
// }
// func (u Role) DeleteRoleById(context context.Context, RoleID int) (sql.Result, error) {
// 	query := "DELETE FROM `kynang` WHERE ID = ?"
// 	result, err := u.sqlDB.ExecContext(context, query, RoleID)
// 	if err != nil {
// 		if err.(*mysql.MySQLError).Number == 1451 {
// 			log.Error(err.Error())
// 			return result, banana.ForenkeyErrol
// 		}
// 		log.Error(err.Error())
// 		return result, banana.SererError
// 	}
// 	return result, nil
// }
