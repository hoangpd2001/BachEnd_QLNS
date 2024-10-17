package insurancecontroller


import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	modelInsurance "BackEnd/mod/model/mode_insurance"
	repository "BackEnd/mod/repository/repo_insurance"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type InsuranceController struct {
	InsuranceRepo repository.InsuranceRepo
	Bind      utils.Bind
}

// ========================================================================================================
func (u *InsuranceController) CreatInsurance(c echo.Context) error {
	req := &modelInsurance.ReqInsurance{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	  if err != nil {
        return c.JSON(http.StatusBadRequest, model.Response{
            StatusCode: http.StatusBadRequest,
            Message:    err.Error(),
            Data:       nil,
        })
    }
	req, ok := validatedReq.(*modelInsurance.ReqInsurance)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	res := modelInsurance.ResInsurance{
		TenBaoHiem: req.TenBaoHiem,
		NhaCungCap: req.NhaCungCap,
		NoiDangKi: req.NoiDangKi,
		TyLePhi: req.TyLePhi,
	}
	InsuranceResult, err := u.InsuranceRepo.CreatInsurance(c.Request().Context(), res)
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
		Data:       InsuranceResult,
	})
}

// //========================================================================================================

func (u *InsuranceController) SelectInsuranceAll(c echo.Context) error {
	InsuranceResult, err := u.InsuranceRepo.SelectInsuranceAll(c.Request().Context())
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
		Data:       InsuranceResult,
	})
}

// //========================================================================================================

func (u *InsuranceController) SelelectInsuranceById(c echo.Context) error {
	idInsurance, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	InsuranceResult, err := u.InsuranceRepo.SelelectInsuranceById(c.Request().Context(), idInsurance)
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
		Data:       InsuranceResult,
	})
}

//========================================================================================================

func (u *InsuranceController) UpdateInsuranceById(c echo.Context) error {
	idInsurance, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	req := &modelInsurance.ReqInsurance{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	  if err != nil {
        return c.JSON(http.StatusBadRequest, model.Response{
            StatusCode: http.StatusBadRequest,
            Message:    err.Error(),
            Data:       nil,
        })
    }
	req, ok := validatedReq.(*modelInsurance.ReqInsurance)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}

	Insurance := modelInsurance.ResInsurance{
		TenBaoHiem: req.TenBaoHiem,
		NhaCungCap: req.NhaCungCap,
		NoiDangKi: req.NoiDangKi,
		TyLePhi: req.TyLePhi,
		ID:           idInsurance,
	}
	InsuranceResult, err := u.InsuranceRepo.UpdateInsuranceById(c.Request().Context(), Insurance)
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
		Data:       InsuranceResult,
	})
}

// ====================================================================================================================
func (u *InsuranceController) DeleteInsuranceById(c echo.Context) error {
	idInsurance, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.InsuranceRepo.DeleteInsuranceById(c.Request().Context(), idInsurance)
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
