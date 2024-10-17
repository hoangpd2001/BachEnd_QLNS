package usercontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	reqUser "BackEnd/mod/model/model_user/req_user"
	resUser "BackEnd/mod/model/model_user/res_user"
	repository "BackEnd/mod/repository/repo_user"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EducationController struct {
	EducationRepo repository.EducationRepo
	Bind          utils.Bind
}

// ========================================================================================================
func (u *EducationController) CreatEducation(c echo.Context) error {
	req := &reqUser.ReqEducation{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req, ok := validatedReq.(*reqUser.ReqEducation)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	UserEdu := resUser.ResEducation{
		IDNhanVien:   req.IDNhanVien,
		Truong:       req.Truong,
		BangCap:      req.BangCap,
		CapHoc:       req.CapHoc,
		NamTotNghiep: req.NamTotNghiep,
	}
	userR, err := u.EducationRepo.CreatEducation(c.Request().Context(), UserEdu)
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

func (u *EducationController) SelectEducationByUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	userR, err := u.EducationRepo.SelectEducationByUser(c.Request().Context(), idUser)
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

func (u *EducationController) SelectEducationById(c echo.Context) error {
	idEducation, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	userR, err := u.EducationRepo.SelectEducationById(c.Request().Context(), idEducation)
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

func (u *EducationController) UpdateEducationById(c echo.Context) error {
	req := &reqUser.ReqEducation{}
	idEducation, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req, ok := validatedReq.(*reqUser.ReqEducation)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	UserEdu := resUser.ResEducation{
		ID:           idEducation,
		Truong:       req.Truong,
		BangCap:      req.BangCap,
		CapHoc:       req.CapHoc,
		NamTotNghiep: req.NamTotNghiep,
	}
	userR, err := u.EducationRepo.UpdateEducationById(c.Request().Context(), UserEdu)
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

func (u *EducationController) DeleteEducationById(c echo.Context) error {
	idEducation, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.EducationRepo.DeleteEducationById(c.Request().Context(), idEducation)
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
