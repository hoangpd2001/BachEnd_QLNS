package usercontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	reqUser "BackEnd/mod/model/model_user/req_user"
	resuser "BackEnd/mod/model/model_user/res_user"
	repouser "BackEnd/mod/repository/repo_user"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
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
	//	hash := security.HashingPasswordFunc(req.Pass)
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

// func (u *UseController) HandleSignIn(c echo.Context) error {
// 	req := req.ReqSignIn{}
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
// 	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, model.Response{
// 			StatusCode: http.StatusUnauthorized,
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}
// 	isTheSame := security.CheckPasswordHashFunc(req.PassWord,user.PassWord )
// 	if !isTheSame {
// 		return c.JSON(http.StatusUnauthorized, model.Response{
// 			StatusCode: http.StatusUnauthorized,
// 			Message:    "sai mật khẩu",
// 			Data:       nil,
// 		})
// 	}
// 	token, err := security.GenToken(user)
// 	if err!=nil{
// 		log.Error(err)
// 		return c.JSON(http.StatusInternalServerError, model.Response{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    err.Error(),
// 			Data:       nil,

// 	})}
// 	user.PassWord = ""
// 	user.Token = token
// 	return c.JSON(http.StatusOK, model.Response{
// 		StatusCode: http.StatusOK,
// 		Message:    "Thành Công",
// 		Data:       user,
// 	})
// }

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
