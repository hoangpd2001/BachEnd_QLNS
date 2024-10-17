package usercontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	reqUser "BackEnd/mod/model/model_user/req_user"
	resUser "BackEnd/mod/model/model_user/res_user"
	repository "BackEnd/mod/repository/repo_user"
	"BackEnd/mod/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RelativeController struct {
	RelativeRepo repository.RelativeRepo
	Bind         utils.Bind
}

// ========================================================================================================
func (u *RelativeController) CreatRelative(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	req := &reqUser.ReqRelative{}

	validatedReq, err := u.Bind.BindAndValidate(c, req)
	  if err != nil {
        return c.JSON(http.StatusBadRequest, model.Response{
            StatusCode: http.StatusBadRequest,
            Message:    err.Error(),
            Data:       nil,
        })
    }
	req, ok := validatedReq.(*reqUser.ReqRelative)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}

	UserRelative := resUser.ResRelative{
		IDNhanVien:      idUser,
		TenNguoiThan:    req.TenNguoiThan,
		SDTNguoiThan:    req.SDTNguoiThan,
		DiaChiNguoiThan: req.DiaChiNguoiThan,
		QuanHe:          req.QuanHe,
	}
	userR, err := u.RelativeRepo.CreatRelative(c.Request().Context(), UserRelative)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Thành Công",
		Data:       userR,
	})
}

//========================================================================================================

func (u *RelativeController) SelectRelativeByUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	userR, err := u.RelativeRepo.SelectRelativeByUser(c.Request().Context(), idUser)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Thành Công",
		Data:       userR,
	})
}

//========================================================================================================

func (u *RelativeController) UpdateRelativeByUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	req := &reqUser.ReqRelative{}
	validatedReq, err := u.Bind.BindAndValidate(c,req)
  if err != nil {
        return c.JSON(http.StatusBadRequest, model.Response{
            StatusCode: http.StatusBadRequest,
            Message:    err.Error(),
            Data:       nil,
        })
    }
	req, ok := validatedReq.(*reqUser.ReqRelative)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}

	UserRelative := resUser.ResRelative{
		IDNhanVien:      idUser,
		TenNguoiThan:    req.TenNguoiThan,
		SDTNguoiThan:    req.SDTNguoiThan,
		DiaChiNguoiThan: req.DiaChiNguoiThan,
		QuanHe:          req.QuanHe,
	}
	userR, err := u.RelativeRepo.UpdateRelativeByUser(c.Request().Context(), UserRelative)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Thành Công",
		Data:       userR,
	})
}
