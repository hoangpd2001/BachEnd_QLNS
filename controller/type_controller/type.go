package typecontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	modetypeuser "BackEnd/mod/model/mode_typeUser"
	repository "BackEnd/mod/repository/repo_type"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TypeController struct {
	TypeRepo repository.TypeRepo
	Bind     utils.Bind
}

// ========================================================================================================
func (u *TypeController) CreatType(c echo.Context) error {
	req := &modetypeuser.ReqTypeUser{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modetypeuser.ReqTypeUser)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	res := modetypeuser.ResTypeUser{
		LoaiNhanVien: req.LoaiNhanVien,
	}
	userR, err := u.TypeRepo.CreatTYpe(c.Request().Context(), res)
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

// //========================================================================================================

func (u *TypeController) SelectTypeAll(c echo.Context) error {
	userR, err := u.TypeRepo.SelectTypeAll(c.Request().Context())
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

// //========================================================================================================

func (u *TypeController) SelelectTypeByUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	userR, err := u.TypeRepo.SelelectTypeByUser(c.Request().Context(), idUser)
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

func (u *TypeController) UpdateTypeById(c echo.Context) error {
	idType, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	req := &modetypeuser.ReqTypeUser{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modetypeuser.ReqTypeUser)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}

	UserType := modetypeuser.ResTypeUser{
		LoaiNhanVien: req.LoaiNhanVien,
		ID:           idType,
	}
	userR, err := u.TypeRepo.UpdateTypeById(c.Request().Context(), UserType)
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

// ====================================================================================================================
func (u *TypeController) DeleteTypeById(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.TypeRepo.DeleteTypeById(c.Request().Context(), idUser)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	row, _ := result.RowsAffected()
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Xóa thành công %d hàng", row),
		Data:       result,
	})
}
