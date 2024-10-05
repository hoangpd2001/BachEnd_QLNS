package titlecontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	repotitle "BackEnd/mod/repository/repo_title"
	modelTitle "BackEnd/mod/model/model_title"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TitleController struct {
	TitleRepo repotitle.TitleRepo
	Bind      utils.Bind
}

// ========================================================================================================
func (u *TitleController) CreatTitle(c echo.Context) error {
	req := &modelTitle.ReqTitle{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modelTitle.ReqTitle)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	res := modelTitle.ResTitle{
		TenChucDanh: req.TenChucDanh,
	}
	userR, err := u.TitleRepo.CreatTitle(c.Request().Context(), res)
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

func (u *TitleController) SelectTitleAll(c echo.Context) error {
	userR, err := u.TitleRepo.SelectTitleAll(c.Request().Context())
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

func (u *TitleController) SelelectTitleById(c echo.Context) error {
	idTitle, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	userR, err := u.TitleRepo.SelelectTitleById(c.Request().Context(), idTitle)
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

func (u *TitleController) UpdateTitleById(c echo.Context) error {
	idTitle, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	req := &modelTitle.ReqTitle{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modelTitle.ReqTitle)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}

	Title := modelTitle.ResTitle{
		TenChucDanh: req.TenChucDanh,
		ID:       idTitle,
	}
	userR, err := u.TitleRepo.UpdateTitleById(c.Request().Context(), Title)
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
func (u *TitleController) DeleteTitleById(c echo.Context) error {
	idTitle, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.TitleRepo.DeleteTitleById(c.Request().Context(), idTitle)
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
