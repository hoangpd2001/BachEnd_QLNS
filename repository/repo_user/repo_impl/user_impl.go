package repoimpl

import (
	"BackEnd/mod/banana"
	reqUser "BackEnd/mod/model/model_user/req_user"
	resUser "BackEnd/mod/model/model_user/res_user"
	repouser "BackEnd/mod/repository/repo_user"
	"context"
	"database/sql"
	"fmt"

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
	 SDT, NgaySinh, DiaChi, CCCD, IDLoaiNhanVien, IDCapBac, NgayBatDau, NgayKetThuc,MatKhau) 
	 VALUES (:Ten,:Dem,:Ho,:Email,:GioiTinh,:SDT,:NgaySinh,:DiaChi,:CCCD,:IDLoaiNhanVien,:IDCapBac,:NgayBatDau,:NgayKetThuc,:MatKhau)
	`
	result, err := u.sqlDB.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1062 {
				return user, banana.UserConflict

			}
		}

		return user, banana.SererError
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
	sql := `SELECT * FROM nhanvien`
	err := u.sqlDB.SelectContext(context, &listUser, sql)
	if err != nil {
		log.Error(err.Error())
		return listUser, err
	}

	return listUser, nil
}

//=================================================================================================================================

func (u UserRepoImpl) SelectUserById(context context.Context, UserId int) (resUser.ResUser, error) {
	var user resUser.ResUser
	sql := `SELECT 
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
	LEFT JOIN nhanvien_nguoithan nguoiThan ON nv.ID = nguoiThan.IDNhanVien,
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

// =====================================================================================================================
func (u UserRepoImpl) SelectCountUser(context context.Context) ([]resUser.ResUserCount, error) {
	var listUser []resUser.ResUserCount
	sql := `
		SELECT 
  DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL seq MONTH), '%Y-%m') AS Thang,
  COUNT(nv.ID) AS SoLuong
FROM 
  (SELECT 0 AS seq UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 
   UNION SELECT 5) AS months
LEFT JOIN nhanvien nv
ON nv.NgayBatDau <= LAST_DAY(DATE_SUB(CURDATE(), INTERVAL seq MONTH))
AND (nv.NgayKetThuc IS NULL OR nv.NgayKetThuc = '0000-00-00' OR nv.NgayKetThuc >= DATE_SUB(CURDATE(), INTERVAL seq MONTH))
GROUP BY Thang
ORDER BY Thang;

	`
	err := u.sqlDB.SelectContext(context, &listUser, sql)
	if err != nil {
		log.Error(err.Error())
		return listUser, err
	}
	return listUser, nil
}

func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq reqUser.ReqSignIn) ([]resUser.ResSingin, error) {
	var ListUser = []resUser.ResSingin{}
	sql := `SELECT 
					n.ID,
					n.Email,
					n.MatKhau,
					COALESCE(nc.IDChucDanh, 0) AS IDChucDanh
				FROM 
					nhanvien n
				LEFT JOIN 
					nhanvien_chucdanh nc 
				ON 
					n.ID = nc.IDNhanVien 
					AND nc.NgayKetThuc = '0000-00-00 00:00:00'
				WHERE 
					n.Email = ?;`
	err := u.sqlDB.SelectContext(context, &ListUser, sql, loginReq.Email)
	if err != nil {
		return ListUser, err
	}

	return ListUser, nil
}

func (u *UserRepoImpl) EditLogin(context context.Context, loginReq reqUser.ReqSignInEdit, mk string) (sql.Result, error) {
	sql := `UPDATE 
				nhanvien
			SET MatKhau = ?
			WHERE 
				Email	 = ?
				AND SDT  = ?
			`
	fmt.Println(mk)
	result, err := u.sqlDB.ExecContext(context, sql, mk, loginReq.Email, loginReq.SDT)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (u UserRepoImpl) UpdateUserById(context context.Context, user resUser.ResUser) (resUser.ResUser, error) {
	statement := `
		UPDATE nhanvien SET Ten=:Ten,Dem=:Dem,Ho=:Ho,
		Email=:Email,GioiTinh=:GioiTinh,SDT=:SDT,
		NgaySinh=:NgaySinh,DiaChi=:DiaChi,CCCD=:CCCD,
		IDLoaiNhanVien=:IDLoaiNhanVien,IDCapBac=:IDCapBac,NgayBatDau=:NgayBatDau,
		NgayKetThuc=:NgayKetThuc WHERE ID =:ID
	`
	_, err := u.sqlDB.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1062 {
				return user, banana.UserConflict
			}
		}

		return user, banana.SererError
	}
	return user, nil
}
