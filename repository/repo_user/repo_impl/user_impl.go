package repoimpl

import (
	"BackEnd/mod/banana"
	resUser "BackEnd/mod/model/model_user/res_user"
	repouser "BackEnd/mod/repository/repo_user"
	"context"
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
	INSERT INTO nhanvien( Ten, Dem, Ho, Email, GioiTinh,
	 SDT, NgaySinh, DiaChi, CCCD, IDLoaiNhanVien, IDCapBac, NgayBatDau, NgayKetThuc) 
	 VALUES (:Ten,:Dem,:Ho,:Email,:GioiTinh,:SDT,:NgaySinh,:DiaChi,:CCCD,:IDLoaiNhanVien,:IDCapBac,:NgayBatDau,:NgayKetThuc)
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
	return user, nil
}

// =====================================================================================================================
func (u UserRepoImpl) SelectUserAll(context context.Context) ([]resUser.ResUser, error) {
	var listUser []resUser.ResUser
	sql:=`SELECT 
    nv.ID, nv.Ten, nv.Ho,nv.Dem, nv.Email, nv.GioiTinh, nv.SDT, nv.NgaySinh, nv.DiaChi, nv.CCCD,nv.NgayBatDau,nv.NgayKetThuc,
    loainv.LoaiNhanVien,
    capbac.TenCapBac,
    chucdanh.IDChucDanh,
    phongban.TenPhongBan,
    chinhanh.ChiNhanh
	FROM 
    nhanvien nv
	LEFT JOIN loainhanvien loainv ON nv.IDLoaiNhanVien = loainv.ID
	LEFT JOIN capbac capbac ON nv.IDCapBac = capbac.ID
	LEFT JOIN nhanvien_chucdanh chucdanh ON nv.ID = chucdanh.IDNhanVien 
    AND chucdanh.NgayKetThuc ="0000-00-00 00:00:00"
    AND chucdanh.NgayBatDau !="0000-00-00 00:00:00"
	LEFT JOIN phongban phongban ON chucdanh.IDPhongBan = phongban.ID
	LEFT JOIN chinhanh chinhanh ON phongban.IDChiNhanh = chinhanh.ID
	`
	err := u.sqlDB.SelectContext(context, &listUser,sql)
	if err != nil {
		log.Error(err.Error())
		return listUser, err
	}

	return listUser, nil
}

//=================================================================================================================================

func (u UserRepoImpl) SelectUserById(context context.Context, UserId int) (resUser.ResUser, error) {
	var user resUser.ResUser
	sql:=`SELECT 
    nv.ID, nv.Ten, nv.Ho,nv.Dem, nv.Email, nv.GioiTinh, nv.SDT, nv.NgaySinh, nv.DiaChi, nv.CCCD,nv.NgayBatDau,nv.NgayKetThuc,
    loainv.LoaiNhanVien,
    capbac.TenCapBac,
    nguoiThan.TenNguoiThan, nguoiThan.SDTNguoiThan, nguoiThan.QuanHe, nguoiThan.DiaChiNguoiThan,
    chucdanh.IDChucDanh, 
    phongban.TenPhongBan,
    chinhanh.ChiNhanh
	FROM 
    nhanvien nv
	LEFT JOIN loainhanvien loainv ON nv.IDLoaiNhanVien = loainv.ID
	LEFT JOIN capbac capbac ON nv.IDCapBac = capbac.ID
	LEFT JOIN nhanvien_nguoithan nguoiThan ON nv.ID = nguoiThan.IDNhanVien
	LEFT JOIN nhanvien_chucdanh chucdanh ON nv.ID = chucdanh.IDNhanVien 
    AND chucdanh.NgayKetThuc ="0000-00-00 00:00:00"
    AND chucdanh.NgayBatDau !="0000-00-00 00:00:00"
	LEFT JOIN phongban phongban ON chucdanh.IDPhongBan = phongban.ID
	LEFT JOIN chinhanh chinhanh ON phongban.IDChiNhanh = chinhanh.ID
	Where nv.ID = ?`
	err := u.sqlDB.GetContext(context, &user, sql, UserId)

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1062 {
				return user, banana.GetIdFailed

			}
		}

		return user, banana.SignUpFail
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
