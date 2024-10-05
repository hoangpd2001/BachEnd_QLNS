package branchcontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	modelbranch "BackEnd/mod/model/model_Branch"
	repobranch "BackEnd/mod/repository/repo_branch"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BranchController struct {
	BranchRepo repobranch.BranchRepo
	Bind     utils.Bind
}

// ========================================================================================================
func (u *BranchController) CreatBranch(c echo.Context) error {
	req := &modelbranch.ReqBranch{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modelbranch.ReqBranch)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	res := modelbranch.ResBranch{
		ChiNhanh: req.ChiNhanh,
	}
	userR, err := u.BranchRepo.CreatBranch(c.Request().Context(), res)
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

func (u *BranchController) SelectBranchAll(c echo.Context) error {
	userR, err := u.BranchRepo.SelectBranchAll(c.Request().Context())
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

func (u *BranchController) SelelectBranchById(c echo.Context) error {
	idBranch, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	userR, err := u.BranchRepo.SelelectBranchById(c.Request().Context(), idBranch)
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

func (u *BranchController) UpdateBranchById(c echo.Context) error {
	idBranch, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	req := &modelbranch.ReqBranch{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modelbranch.ReqBranch)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}

	Branch := modelbranch.ResBranch{
		ChiNhanh: req.ChiNhanh,
		ID:           idBranch,
	}
	userR, err := u.BranchRepo.UpdateTypeById(c.Request().Context(), Branch)
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
func (u *BranchController) DeleteBranchById(c echo.Context) error {
	idBranch, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.BranchRepo.DeleteBranchById(c.Request().Context(), idBranch)
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
