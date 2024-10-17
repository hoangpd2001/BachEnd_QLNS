package grupcontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	modelGrup "BackEnd/mod/model/model_grup"
	repogrup "BackEnd/mod/repository/repo_grup"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GrupController struct {
	GrupRepo repogrup.GrupRepo
	Bind     utils.Bind
}

// ========================================================================================================
func (u *GrupController) CreatGrup(c echo.Context) error {
	req := &modelGrup.ReqGrup{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req, ok := validatedReq.(*modelGrup.ReqGrup)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	res := modelGrup.ResGrup{
		TenNhom: req.TenNhom,
	}
	GrupResult, err := u.GrupRepo.CreatGrup(c.Request().Context(), res)
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
		Data:       GrupResult,
	})
}

// //========================================================================================================

func (u *GrupController) SelectGrupAll(c echo.Context) error {
	GrupResult, err := u.GrupRepo.SelectGrupAll(c.Request().Context())
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
		Data:       GrupResult,
	})
}

// //========================================================================================================

//========================================================================================================

func (u *GrupController) UpdateGrupById(c echo.Context) error {
	idGrup, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	req := &modelGrup.ReqGrup{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req, ok := validatedReq.(*modelGrup.ReqGrup)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}

	Grup := modelGrup.ResGrup{
		TenNhom: req.TenNhom,
		ID:      idGrup,
	}
	GrupResult, err := u.GrupRepo.UpdateGrup(c.Request().Context(), Grup)
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
		Data:       GrupResult,
	})
}

// ====================================================================================================================
func (u *GrupController) DeleteGrupById(c echo.Context) error {
	idGrup, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.GrupRepo.DeleteGrup(c.Request().Context(), idGrup)
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
