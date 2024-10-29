package rolecontroller

import (
//	"BackEnd/mod/banana"
	"BackEnd/mod/model"
//	modelRole "BackEnd/mod/model/model_Role"
	repository "BackEnd/mod/repository/repo_Role"
	"BackEnd/mod/utils"

	"net/http"
//	"strconv"

	"github.com/labstack/echo/v4"
)

type RoleController struct {
	RoleRepo repository.RoleRepo
	Bind      utils.Bind
}

// ========================================================================================================
// func (u *RoleController) CreatRole(c echo.Context) error {
// 	req := &modelRole.ReqRole{}
// 	validatedReq, err := u.Bind.BindAndValidate(c, req)
// 	  if err != nil {
//         return c.JSON(http.StatusBadRequest, model.Response{
//             StatusCode: http.StatusBadRequest,
//             Message:    err.Error(),
//             Data:       nil,
//         })
//     }
// 	req, ok := validatedReq.(*modelRole.ReqRole)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, model.Response{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    "Failed to cast validated request",
// 			Data:       nil,
// 		})
// 	}
// 	res := modelRole.ResRole{
// 		TenKyNang: req.TenKyNang,
// 		MoTa: req.MoTa,
// 	}
// 	RoleResult, err := u.RoleRepo.CreatRole(c.Request().Context(), res)
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
// 		Data:       RoleResult,
// 	})
// }

// //========================================================================================================

func (u *RoleController) SelectRoleAll(c echo.Context) error {
	RoleResult, err := u.RoleRepo.SelectRoleAll(c.Request().Context())
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
		Data:       RoleResult,
	})
}

// //========================================================================================================

// func (u *RoleController) SelelectRoleById(c echo.Context) error {
// 	idRole, err := strconv.Atoi(c.QueryParam("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusConflict, model.Response{
// 			StatusCode: http.StatusConflict,
// 			Message:    banana.GetIdFailed.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	RoleResult, err := u.RoleRepo.SelelectRoleById(c.Request().Context(), idRole)
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
// 		Data:       RoleResult,
// 	})
// }

// //========================================================================================================

// func (u *RoleController) UpdateRoleById(c echo.Context) error {
// 	idRole, err := strconv.Atoi(c.QueryParam("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusConflict, model.Response{
// 			StatusCode: http.StatusConflict,
// 			Message:    banana.GetIdFailed.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	req := &modelRole.ReqRole{}
// 	validatedReq, err := u.Bind.BindAndValidate(c, req)
// 	  if err != nil {
//         return c.JSON(http.StatusBadRequest, model.Response{
//             StatusCode: http.StatusBadRequest,
//             Message:    err.Error(),
//             Data:       nil,
//         })
//     }
// 	req, ok := validatedReq.(*modelRole.ReqRole)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, model.Response{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    "Failed to cast validated request",
// 			Data:       nil,
// 		})
// 	}

// 	Role := modelRole.ResRole{
// 		TenKyNang: req.TenKyNang,
// 		MoTa: req.MoTa,
// 		ID:           idRole,
// 	}
// 	RoleResult, err := u.RoleRepo.UpdateRoleById(c.Request().Context(), Role)
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
// 		Data:       RoleResult,
// 	})
// }

// // ====================================================================================================================
// func (u *RoleController) DeleteRoleById(c echo.Context) error {
// 	idRole, err := strconv.Atoi(c.QueryParam("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusConflict, model.Response{
// 			StatusCode: http.StatusConflict,
// 			Message:    banana.GetIdFailed.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	result, err := u.RoleRepo.DeleteRoleById(c.Request().Context(), idRole)
// 	if err != nil {
// 		return c.JSON(http.StatusConflict, model.Response{
// 			StatusCode: http.StatusConflict,
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}
// 	row, _ := result.RowsAffected()
// 	return c.JSON(http.StatusOK, model.Response{
// 		StatusCode: http.StatusOK,
// 		Message:    fmt.Sprintf("Xóa thành công %d hàng", row),
// 		Data:       result,
// 	})
// }
