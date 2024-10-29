package rolecontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	modelrole "BackEnd/mod/model/model_role"
	repository "BackEnd/mod/repository/repo_Role"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoleUserController struct {
	RoleUserRepo repository.RoleUserRepo
	Bind         utils.Bind
	CustomDate   utils.CustomDate
}

// ========================================================================================================
func (u *RoleUserController) CreatRoleUser(c echo.Context) error {
	idRole, err := strconv.Atoi(c.QueryParam("idr"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	idTitle, err := strconv.Atoi(c.QueryParam("idt"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	req := &modelrole.ReqUserRole{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req, ok := validatedReq.(*modelrole.ReqUserRole)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	res := modelrole.ResUserRole{
		IDChucDanh: idTitle,
		IDVaiTro:   idRole,
		Xem:        req.Xem,
		Them:       req.Them,
		Sua:        req.Sua,
		Xoa:        req.Xoa,
	}
	skillResult, err := u.RoleUserRepo.CreatRoleUser(c.Request().Context(), res)
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

func (u *RoleUserController) SelectRoleUserAll(c echo.Context) error {
	skillResult, err := u.RoleUserRepo.SelectRoleUserAll(c.Request().Context())
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

// func (u *RoleUserController) SelectRoleUser(c echo.Context) error {

// 	skillResult, err := u.RoleUserRepo.SelectRoleUser(c.Request().Context())
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
// 		Data:       skillResult,
// 	})
// }

// // //========================================================================================================

func (u *RoleUserController) SelelectRoleUser(c echo.Context,idTitle []int, idRole int) error {
	skillResult, err := u.RoleUserRepo.SelelectRoleUser(c.Request().Context(), idTitle, idRole)
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

func (u *RoleUserController) UpdateRoleUser(c echo.Context) error {
	idRole, err := strconv.Atoi(c.QueryParam("idr"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	idTitle, err := strconv.Atoi(c.QueryParam("idt"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	req := &modelrole.ReqUserRole{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req, ok := validatedReq.(*modelrole.ReqUserRole)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	skill := modelrole.ResUserRole{
		IDChucDanh: idTitle,
		IDVaiTro:   idRole,
		Xem:        req.Xem,
		Them:       req.Them,
		Sua:        req.Sua,
		Xoa:        req.Xoa,
	}
	skillResult, err := u.RoleUserRepo.UpdateSkillById(c.Request().Context(), skill)
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
func (u *RoleUserController) DeleteRoleUser(c echo.Context) error {
	idRole, err := strconv.Atoi(c.QueryParam("idr"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	idTitle, err := strconv.Atoi(c.QueryParam("idt"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.RoleUserRepo.DeleteSkillById(c.Request().Context(), idTitle, idRole)
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
