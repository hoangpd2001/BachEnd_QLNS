package levelcontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	modelevel "BackEnd/mod/model/mode_level"
	repolevel "BackEnd/mod/repository/repo_level"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LevelController struct {
	LevelRepo repolevel.LevelRepo
	Bind      utils.Bind
}

// ========================================================================================================
func (u *LevelController) CreatLevel(c echo.Context) error {
	req := &modelevel.ReqLevel{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modelevel.ReqLevel)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	res := modelevel.ResLevel{
		TenCapBac: req.TenCapBac,
		CauTrucLuong: req.CauTrucLuong,
	}
	level, err := u.LevelRepo.CreatLevel(c.Request().Context(), res)
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
		Data:       level,
	})
}

// //========================================================================================================

func (u *LevelController) SelectLevelAll(c echo.Context) error {
	level, err := u.LevelRepo.SelectLevelAll(c.Request().Context())
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
		Data:       level,
	})
}

// //========================================================================================================

func (u *LevelController) SelelectLevelByUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	level, err := u.LevelRepo.SelelectLevelByUser(c.Request().Context(), idUser)
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
		Data:       level,
	})
}

//========================================================================================================

func (u *LevelController) UpdateLevelById(c echo.Context) error {
	idLevel, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	req := &modelevel.ReqLevel{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modelevel.ReqLevel)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}

	UserLevel := modelevel.ResLevel{
		TenCapBac: req.TenCapBac,
		CauTrucLuong: req.CauTrucLuong,
		ID:           idLevel,
	}
	level, err := u.LevelRepo.UpdateLevelById(c.Request().Context(), UserLevel)
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
		Data:       level,
	})
}

// ====================================================================================================================
func (u *LevelController) DeleteLevelById(c echo.Context) error {
	idLevel, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.LevelRepo.DeleteLevelById(c.Request().Context(), idLevel)
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
