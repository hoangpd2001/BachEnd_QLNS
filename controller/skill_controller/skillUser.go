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

type SkillUserController struct {
	SkillUserRepo repository.SkillUserRepo
	Bind          utils.Bind
	CustomDate    utils.CustomDate
}

// ========================================================================================================
func (u *SkillUserController) CreatSkillUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	idSkill, err := strconv.Atoi(c.QueryParam("ids"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	req := &modelskill.ReqUserSkill{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modelskill.ReqUserSkill)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	err = u.CustomDate.UnmarshalJSON([]byte(req.NgayDanhGia))
	res := modelskill.ResUserSkill{
		IDNhanVien:  idUser,
		IDKyNang:    idSkill,
		MucDo:       req.MucDo,
		NgayDanhGia: u.CustomDate.Time,
	}
	skillResult, err := u.SkillUserRepo.CreatSkillUser(c.Request().Context(), res)
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

func (u *SkillUserController) SelectSkillUserAll(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	skillResult, err := u.SkillUserRepo.SelectSkillUserAll(c.Request().Context(), idUser)
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

func (u *SkillUserController) SelelectSkillUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	idSkill, err := strconv.Atoi(c.QueryParam("ids"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	skillResult, err := u.SkillUserRepo.SelelectSkillUser(c.Request().Context(), idSkill, idUser)
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

func (u *SkillUserController) UpdateSkillUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	idSkill, err := strconv.Atoi(c.QueryParam("ids"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	req := &modelskill.ReqUserSkill{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modelskill.ReqUserSkill)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	err = u.CustomDate.UnmarshalJSON([]byte(req.NgayDanhGia))
	skill := modelskill.ResUserSkill{
		IDNhanVien:  idUser,
		IDKyNang:    idSkill,
		MucDo:       req.MucDo,
		NgayDanhGia: u.CustomDate.Time,
	}
	skillResult, err := u.SkillUserRepo.UpdateSkillById(c.Request().Context(), skill)
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
func (u *SkillUserController) DeleteSkillUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	idSkill, err := strconv.Atoi(c.QueryParam("ids"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.SkillUserRepo.DeleteSkillById(c.Request().Context(), idSkill, idUser)
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
