package usercontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	reqUser "BackEnd/mod/model/model_user/req_user"
	resuser "BackEnd/mod/model/model_user/res_user"
	repouser "BackEnd/mod/repository/repo_user"
	"BackEnd/mod/security"
	"BackEnd/mod/utils"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UseController struct {
	UserRepo   repouser.UserRepo
	CustomDate utils.CustomDate
	Bind       utils.Bind
}

//==============================================================================================================

func (u *UseController) CreatUser(c echo.Context) error {
	req := &reqUser.ReqUser{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req, ok := validatedReq.(*reqUser.ReqUser)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	hash := security.HashingPasswordFunc(req.Pass)
	//	role := model.MEMBER.String()
	err = u.CustomDate.UnmarshalJSON([]byte(`"` + req.NgaySinh + `"`))
	if err != nil {
		fmt.Println("Error parsing NgayBatDau:", err)
	}
	user := resuser.ResUser{

		Ten:            req.Ten,
		Dem:            req.Dem,
		Ho:             req.Ho,
		Email:          req.Email,
		NgaySinh:       u.CustomDate.Time,
		GioiTinh:       req.GioiTinh,
		SDT:            req.SDT,
		DiaChi:         req.DiaChi,
		CCCD:           req.CCCD,
		IDLoaiNhanVien: req.IDLoaiNhanVien,
		IDCapBac:       req.IDCapBac,
		NgayBatDau:     time.Now(),
		MatKhau:        hash,
	}
	userR, err := u.UserRepo.CreatUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	// token, err := security.GenTokenUserFull(user)
	// if err != nil {
	// 	log.Error(err)
	// 	return c.JSON(http.StatusInternalServerError, model.Response{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    err.Error(),
	// 		Data:       nil,
	// 	})
	// }
	// userR.Token = token
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Thành Công",
		Data:       userR,
	})
}

//==============================================================================================================

func (u *UseController) SelectUserAll(c echo.Context) error {
	// tokenData, ok:= c.Get("user").(*jwt.Token)
	// if !ok {
	// 	return c.JSON(http.StatusNotFound, model.Response{
	// 			StatusCode: http.StatusNotFound,
	// 			Message:    banana.NotSignIn.Error(),
	// 			Data:       nil,
	// 		})
	// }
	// claims,ok := tokenData.Claims.(*model.JwtCustomClaims)
	// 	if !ok {
	// 	return c.JSON(http.StatusNotFound, model.Response{
	// 			StatusCode: http.StatusNotFound,
	// 			Message:    banana.NotSignIn.Error(),
	// 			Data:       nil,
	// 		})
	// }
	listUser, err := u.UserRepo.SelectUserAll(c.Request().Context())
	if err != nil {
		if err == banana.UserNotFound {
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       listUser,
	})
}

// ==============================================================================================================
func (u *UseController) SelectUserById(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))

	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	User, err := u.UserRepo.SelectUserById(c.Request().Context(), idUser)
	if err != nil {
		if err == banana.UserNotFound {
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       User,
			})
		}
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       User,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       User,
	})
}
func (u *UseController) SelectCountUser(c echo.Context) error {

	listCount, err := u.UserRepo.SelectCountUser(c.Request().Context())
	if err != nil {
		if err == banana.UserNotFound {
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       listCount,
			})
		}
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       listCount,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       listCount,
	})
}

// func (u *UseController) UpdateUserById(c echo.Context) error{
// 	req := req.ReqUpdateUser{}
// 	if err := c.Bind(&req); err != nil {
// 		log.Error(err.Error())
// 		return c.JSON(http.StatusBadRequest, model.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	validato := validator.New()
// 	if err := validato.Struct(req); err != nil {
// 		log.Error(err.Error())
// 		return c.JSON(http.StatusBadRequest, model.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}
// 	tokenData, ok:= c.Get("user").(*jwt.Token)
// 	if !ok {
// 		return c.JSON(http.StatusNotFound, model.Response{
// 				StatusCode: http.StatusNotFound,
// 				Message:    banana.NotSignIn.Error(),
// 				Data:       nil,
// 			})
// 	}
// 	claims,ok := tokenData.Claims.(*model.JwtCustomClaims)
// 		if !ok {
// 		return c.JSON(http.StatusNotFound, model.Response{
// 				StatusCode: http.StatusNotFound,
// 				Message:    banana.NotSignIn.Error(),
// 				Data:       nil,
// 			})
// 	}
// 	user := model.User{
// 			UserId: claims.UserId,
// 			FullName: req.FullName,
// 			Email:    req.Email,
// 			Role:     req.Role,
// 		}
// 	userR, err := u.UserRepo.UpdateUserById(c.Request().Context(), user)
// 	if err != nil {
// 		return c.JSON(http.StatusConflict, model.Response{
// 			StatusCode: http.StatusConflict,
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}
// 	return c.JSON(http.StatusOK, model.Response{
// 		StatusCode: http.StatusOK,
// 		Message:    "Thành Công",
// 		Data:       userR,
// 	})
// }

//==============================================================================================================

func (u *UseController) HandleSignIn(c echo.Context) error {
	req := &reqUser.ReqSignIn{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req, ok := validatedReq.(*reqUser.ReqSignIn)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Dữ liệu không chính xác",
			Data:       nil,
		})
	}
	ListUser, err := u.UserRepo.CheckLogin(c.Request().Context(), *req)
	fmt.Println(3)
	fmt.Println(ListUser)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    banana.SignInFail.Error(),
			Data:       nil,
		})
	}
	fmt.Println(1)
	fmt.Println(ListUser)
	if len(ListUser) == 0 || ListUser == nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    banana.SignInFail.Error(),
			Data:       nil,
		})
	}
	user := resuser.ResSingin{
		MatKhau: ListUser[0].MatKhau,
		Email:   ListUser[0].Email,
		ID:      ListUser[0].ID,
	}
	fmt.Println(2)
	fmt.Println(ListUser)
	var roles []int
	for _, u := range ListUser {
		if u.Role != 0 {
			roles = append(roles, u.Role)
		} else {
			roles = append(roles, 0)
		}
	}

	isTheSame := security.CheckPasswordHashFunc(req.Pass, user.MatKhau)
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    banana.SignInFail.Error(),
			Data:       nil,
		})
	}
	token, err := security.GenToken(user, roles)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.MatKhau = ""
	user.Token = token
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Thành Công",
		Data:       user,
	})
}

func (u *UseController) HanEditLogin(c echo.Context) error {
	req := &reqUser.ReqSignInEdit{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req, ok := validatedReq.(*reqUser.ReqSignInEdit)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Dữ liệu không chính xác",
			Data:       nil,
		})
	}
	rand.Seed(time.Now().UnixNano())
	max := 99999
	min := 10000
	pass := strconv.Itoa(rand.Intn(max-min+1) + min)
	hash := security.HashingPasswordFunc(pass)
	fmt.Println(hash)
	result, err := u.UserRepo.EditLogin(c.Request().Context(), *req, hash)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    banana.EditLoginFail.Error(),
			Data:       nil,
		})
	}
	count, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    banana.EditLoginFail.Error(),
			Data:       nil,
		})
	}
	if count == 0 {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    banana.EditLoginFail.Error(),
			Data:       nil,
		})
	}
	er := sendEmail(req.Email, "ERP Cap nhat mat khau", "mat khau moi la :"+pass)

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Thành Công",
		Data:       er,
	})
}

// //==============================================================================================================

// func (u *UseController) Profile(c echo.Context) error{
// 	tokenData, ok:= c.Get("user").(*jwt.Token)
// 	if !ok {
// 		return c.JSON(http.StatusNotFound, model.Response{
// 				StatusCode: http.StatusNotFound,
// 				Message:    banana.NotSignIn.Error(),
// 				Data:       nil,
// 			})
// 	}
// 	claims,ok := tokenData.Claims.(*model.JwtCustomClaims)
// 		if !ok {
// 		return c.JSON(http.StatusNotFound, model.Response{
// 				StatusCode: http.StatusNotFound,
// 				Message:    banana.NotSignIn.Error(),
// 				Data:       nil,
// 			})
// 	}
// 	user, err := u.UserRepo.SelectUserById(c.Request().Context(), claims.UserId)
// 	if err != nil {
// 		if err == banana.UserNotFound {
// 			return c.JSON(http.StatusNotFound, model.Response{
// 				StatusCode: http.StatusNotFound,
// 				Message:    err.Error(),
// 				Data:       nil,
// 			})
// 		}

// 		return c.JSON(http.StatusInternalServerError, model.Response{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	return c.JSON(http.StatusOK, model.Response{
// 		StatusCode: http.StatusOK,
// 		Message:    "Xử lý thành công",
// 		Data:       user,
// 	})
// }

// //==============================================================================================================
// func (u *UseController) UpdateProfile(c echo.Context) error{
// 	req := req.ReqUpdateUser{}
// 	if err := c.Bind(&req); err != nil {
// 		log.Error(err.Error())
// 		return c.JSON(http.StatusBadRequest, model.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

//		validato := validator.New()
//		if err := validato.Struct(req); err != nil {
//			log.Error(err.Error())
//			return c.JSON(http.StatusBadRequest, model.Response{
//				StatusCode: http.StatusBadRequest,
//				Message:    err.Error(),
//				Data:       nil,
//			})
//		}
//		tokenData, ok:= c.Get("user").(*jwt.Token)
//		if !ok {
//			return c.JSON(http.StatusNotFound, model.Response{
//					StatusCode: http.StatusNotFound,
//					Message:    banana.NotSignIn.Error(),
//					Data:       nil,
//				})
//		}
//		claims,ok := tokenData.Claims.(*model.JwtCustomClaims)
//			if !ok {
//			return c.JSON(http.StatusNotFound, model.Response{
//					StatusCode: http.StatusNotFound,
//					Message:    banana.NotSignIn.Error(),
//					Data:       nil,
//				})
//		}
//		user := model.User{
//				UserId: claims.UserId,
//				FullName: req.FullName,
//				Email:    req.Email,
//				Role:     req.Role,
//			}
//		userR, err := u.UserRepo.UpdateUserById(c.Request().Context(), user)
//		if err != nil {
//			return c.JSON(http.StatusConflict, model.Response{
//				StatusCode: http.StatusConflict,
//				Message:    err.Error(),
//				Data:       nil,
//			})
//		}
//		return c.JSON(http.StatusOK, model.Response{
//			StatusCode: http.StatusOK,
//			Message:    "Thành Công",
//			Data:       userR,
//		})
//	}
func sendEmail(to string, subject string, body string) error {
	// Cấu hình SMTP
	smtpHost :=  os.Getenv("SMTP_HOST")
	smtpPort :=  os.Getenv("SMTP_POST") 
	authEmail :=  os.Getenv("SMTP_EMAIL") 
	authPassword :=  os.Getenv("SMTP_PASS") 

	// Thiết lập nội dung email
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Xác thực
	auth := smtp.PlainAuth("", authEmail, authPassword, smtpHost)

	// Gửi email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, authEmail, []string{to}, msg)
	if err != nil {
		return err
	}
	return nil
}
