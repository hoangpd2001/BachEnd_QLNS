package skillcontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	modelskill "BackEnd/mod/model/model_skill"
	repository "BackEnd/mod/repository/repo_skill"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SkillController struct {
	SkillRepo repository.SkillRepo
	Bind      utils.Bind
}

// ========================================================================================================
func (u *SkillController) CreatSkill(c echo.Context) error {
	req := &modelskill.ReqSkill{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	  if err != nil {
        return c.JSON(http.StatusBadRequest, model.Response{
            StatusCode: http.StatusBadRequest,
            Message:    err.Error(),
            Data:       nil,
        })
    }
	req, ok := validatedReq.(*modelskill.ReqSkill)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	res := modelskill.ResSkill{
		TenKyNang: req.TenKyNang,
		MoTa: req.MoTa,
	}
	skillResult, err := u.SkillRepo.CreatSkill(c.Request().Context(), res)
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
		Data:       skillResult,
	})
}

// //========================================================================================================

func (u *SkillController) SelectSkillAll(c echo.Context) error {
	skillResult, err := u.SkillRepo.SelectSkillAll(c.Request().Context())
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
		Data:       skillResult,
	})
}

// //========================================================================================================

func (u *SkillController) SelelectSkillById(c echo.Context) error {
	idSkill, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	skillResult, err := u.SkillRepo.SelelectSkillById(c.Request().Context(), idSkill)
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
		Data:       skillResult,
	})
}

//========================================================================================================

func (u *SkillController) UpdateSkillById(c echo.Context) error {
	idSkill, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	req := &modelskill.ReqSkill{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	  if err != nil {
        return c.JSON(http.StatusBadRequest, model.Response{
            StatusCode: http.StatusBadRequest,
            Message:    err.Error(),
            Data:       nil,
        })
    }
	req, ok := validatedReq.(*modelskill.ReqSkill)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}

	skill := modelskill.ResSkill{
		TenKyNang: req.TenKyNang,
		MoTa: req.MoTa,
		ID:           idSkill,
	}
	skillResult, err := u.SkillRepo.UpdateSkillById(c.Request().Context(), skill)
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
		Data:       skillResult,
	})
}

// ====================================================================================================================
func (u *SkillController) DeleteSkillById(c echo.Context) error {
	idSkill, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.SkillRepo.DeleteSkillById(c.Request().Context(), idSkill)
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
