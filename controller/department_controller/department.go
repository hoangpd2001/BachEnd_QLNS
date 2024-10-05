package departmentcontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	modeldepartment "BackEnd/mod/model/model_department"
	repodepartment "BackEnd/mod/repository/repo_department"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DepartmentController struct {
	DepartmentRepo repodepartment.Department
	Bind           utils.Bind
	CustomDate     utils.CustomDate
}

// ========================================================================================================
func (u *DepartmentController) CreatDepartment(c echo.Context) error {

	req := &modeldepartment.Reqdepartment{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modeldepartment.Reqdepartment)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}

	res := modeldepartment.ResDepartment{
		TenPhongBan: req.TenPhongBan,
		IDChiNhanh: req.IDChiNhanh,
	}
	skillResult, err := u.DepartmentRepo.CreatDepartment(c.Request().Context(), res)
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

func (u *DepartmentController) SelectDepartmentAll(c echo.Context) error {
	Result, err := u.DepartmentRepo.SelectDepartmentAll(c.Request().Context())
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
		Data:       Result,
	})
}

// //========================================================================================================

func (u *DepartmentController) SelelectDepartmentById(c echo.Context) error {
	idDepartment, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}


	skillResult, err := u.DepartmentRepo.SelelectDepartmentById(c.Request().Context(), idDepartment)
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

func (u *DepartmentController) UpdateDepartmentById(c echo.Context) error {
	
	req := &modeldepartment.Reqdepartment{}
	idDepartment, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modeldepartment.Reqdepartment)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
//	err = u.CustomDate.UnmarshalJSON([]byte(req.NgayDanhGia))
	department := modeldepartment.ResDepartment{
		ID: idDepartment,
		TenPhongBan: req.TenPhongBan,
		IDChiNhanh: req.IDChiNhanh,
	}
	skillResult, err := u.DepartmentRepo.UpdateSkillById(c.Request().Context(), department)
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
func (u *DepartmentController) DeleteDepartment(c echo.Context) error {
	idDepartment, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	result, err := u.DepartmentRepo.DeleteSkillById(c.Request().Context(), idDepartment)
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
