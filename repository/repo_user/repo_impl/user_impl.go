package repoimpl

import (
	"BackEnd/mod/banana"
	resUser "BackEnd/mod/model/model_user/res_user"
	repouser "BackEnd/mod/repository/repo_user"
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type UserRepoImpl struct {
	sqlDB *sqlx.DB
}

func NewUserRepo(sql *sqlx.DB) repouser.UserRepo {
	return &UserRepoImpl{
		sqlDB: sql,
	}
}

func (u UserRepoImpl) CreatUser(context context.Context, user resUser.ResUser) (resUser.ResUser, error) {
	statement :=
		`
	INSERT INTO nhanvien( Ten, Dem, Ho, Email, Pass,
	 		LoaiNhanVien, CapBac, ChiNhanh, NgayBatDau)
	VALUES (:Ten, :Dem, :Ho, :Email, :Pass,
			:LoaiNhanVien, :CapBac, :ChiNhanh, :NgayBatDau)
	`
	result, err := u.sqlDB.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1062 {
				return user, banana.UserConflict

			}
		}

		return user, banana.SignUpFail
	}
	userID, err := result.LastInsertId()
	if err != nil {
		log.Error(err.Error())
		return user, err
	}
	user.ID = int(userID)
	// user.IDNhanVien = user.ID
	// statement =
	// 	`
	// INSERT INTO nhanvien_thongtin(IDNhanVien, GioiTinh, SDT, EmailCaNhan,
	// DiaChiThuongTru, DiaChiTamTru, CCCD) VALUES
	// (:IDNhanVien, :GioiTinh, :SDT, :EmailCaNhan,
	// :DiaChiThuongTru, :DiaChiTamTru, :CCCD)
	// `
	// _, err = u.sqlDB.NamedExecContext(context, statement, user)
	// if err != nil {
	// 	log.Error(err.Error())
	// 	if err, ok := err.(*mysql.MySQLError); ok {
	// 		if err.Number == 1062 {
	// 			return user, banana.UserConflict

	// 		}
	// 	}

	// 	return user, banana.SignUpFail
	// }
	// statement =
	// 	`
	// INSERT INTO nhanvien_nguoithan(IDNhanVien, TenNguoiThan,
	// SDTNguoiThan, QuanHe, DiaChi) VALUES
	// (:IDNhanVien, :TenNguoiThan, :SDTNguoiThan, :QuanHe, :DiaChi)
	// `
	// _, err = u.sqlDB.NamedExecContext(context, statement, user)
	// if err != nil {
	// 	log.Error(err.Error())
	// 	if err, ok := err.(*mysql.MySQLError); ok {
	// 		if err.Number == 1062 {
	// 			return user, banana.UserConflict

	// 		}
	// 	}

	// 	return user, banana.SignUpFail
	// }
	return user, nil
}

// =====================================================================================================================
func (u UserRepoImpl) SelectUserAll(context context.Context) ([]resUser.ResUser, error) {
	var listUser []resUser.ResUser

	err := u.sqlDB.SelectContext(context, &listUser,
		"SELECT * FROM nhanvien ")

	if err != nil {
		if err == sql.ErrNoRows {
			return listUser, banana.UserNotFound
		}
		log.Error(err.Error())
		return listUser, err
	}

	return listUser, nil
}

//=================================================================================================================================

func (u UserRepoImpl) SelectUserById(context context.Context, UserId int) (resUser.ResUserFull, error) {
	var user resUser.ResUserFull
	err := u.sqlDB.GetContext(context, &user, "SELECT * FROM nhanvien Where ID = ?", UserId)

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1062 {
				return user, banana.GetIdFailed

			}
		}

		return user, banana.SignUpFail
	}
	err = u.sqlDB.GetContext(context, &user, "SELECT * FROM nhanvien_nguoithan Where IDNhanVien = ?", UserId)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1062 {
				return user, banana.GetIdFailed

			}
		}

		return user, banana.SignUpFail
	}

	err = u.sqlDB.GetContext(context, &user, "SELECT * FROM nhanvien_thongtin Where IDNhanVien = ?", UserId)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1062 {
				return user, banana.GetIdFailed

			}
		}

		return user, banana.GetIdFailed
	}
	return user, nil
}

// func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error) {
// 	var user = model.User{}
// 	err := u.sqlDB.GetContext(context, &user, "SELECT * FROM users WHERE email=?", loginReq.Email)
// 	if err != nil {
// 		log.Error(err.Error())
// 		if err == sql.ErrNoRows {
// 			return user, banana.UserNotFound
// 		}
// 		return user, err
// 	}
// 	return user, nil
// }
// func (u UserRepoImpl) SelectUserById(context context.Context, userId int) (model.User, error) {
// 	var user model.User

// 	err := u.sqlDB.GetContext(context, &user,
// 		"SELECT * FROM users WHERE ID = ?", userId)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return user, banana.UserNotFound
// 		}
// 		log.Error(err.Error())
// 		return user, err
// 	}

// 	return user, nil
// }

// func (u UserRepoImpl) UpdateUserById(context context.Context, user model.User) (model.User, error) {
// 	statement := `
// 		UPDATE users
// 			SET email = :email, full_name = :full_name, role= :role
// 			WHERE ID = :ID;
// 	`
// 	_, err := u.sqlDB.NamedExecContext(context, statement, user)
// 	if err != nil {
// 		log.Error(err.Error())
// 		if err, ok := err.(*mysql.MySQLError); ok {
// 			if err.Number == 1062 {
// 				return user, banana.UserConflict
// 			}
// 		}

// 		return user, banana.SignUpFail
// 	}
// 	return user, nil
// }
