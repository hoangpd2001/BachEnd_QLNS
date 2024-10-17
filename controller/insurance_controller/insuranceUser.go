package insurancecontroller


import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	modelinsurance "BackEnd/mod/model/mode_insurance"
	repository "BackEnd/mod/repository/repo_insurance"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type InsuranceUserController struct {
	InsuranceUserRepo repository.InsuranceUserRepo
	Bind          utils.Bind
	CustomDate    utils.CustomDate
	CustomDate2    utils.CustomDate
}

// ========================================================================================================
func (u *InsuranceUserController) CreatInsuranceUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	idInsurance, err := strconv.Atoi(c.QueryParam("ids"))
	if err!= nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	req := &modelinsurance.ReqUserInsurance{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
  if err != nil {
        return c.JSON(http.StatusBadRequest, model.Response{
            StatusCode: http.StatusBadRequest,
            Message:    err.Error(),
            Data:       nil,
        })
    }
	req, ok := validatedReq.(*modelinsurance.ReqUserInsurance)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	err = u.CustomDate.UnmarshalJSON([]byte(`"` + req.NgayDong + `"`))
	if err != nil {
		fmt.Println("Error parsing NgayBatDau:", err)
	}
	err = u.CustomDate2.UnmarshalJSON([]byte(`"` + req.NgayHetHan + `"`))
	if err != nil {
		fmt.Println("Error parsing NgayBatDau:", err)
	}
	res := modelinsurance.ResUserInsurance{
		IDNhanVien:  idUser,
		IDBaoHiem:    idInsurance,
		NgayDong:    u.CustomDate.Time ,
		NgayHetHan: u.CustomDate2.Time,
	}
	insuranceResult, err := u.InsuranceUserRepo.CreatInsuranceUser(c.Request().Context(), res)
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
		Data:       insuranceResult,
	})
}

// //========================================================================================================

func (u *InsuranceUserController) SelectInsuranceUserAll(c echo.Context) error {
	insuranceResult, err := u.InsuranceUserRepo.SelectInsuranceUserAll(c.Request().Context())
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
		Data:       insuranceResult,
	})
}

func (u *InsuranceUserController) SelectInsuranceUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	insuranceResult, err := u.InsuranceUserRepo.SelectInsuranceUserByUser(c.Request().Context(),idUser)
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
		Data:       insuranceResult,
	})
}

// //========================================================================================================

func (u *InsuranceUserController) SelelectInsuranceUserOne(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	idInsurance, err := strconv.Atoi(c.QueryParam("ids"))
	if err!= nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	insuranceResult, err := u.InsuranceUserRepo.SelelectInsuranceUserByOne(c.Request().Context(), idInsurance, idUser)
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
		Data:       insuranceResult,
	})
}

//========================================================================================================

func (u *InsuranceUserController) UpdateInsuranceUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	idInsurance, err := strconv.Atoi(c.QueryParam("ids"))
	if err!= nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	req := &modelinsurance.ReqUserInsurance{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	  if err != nil {
        return c.JSON(http.StatusBadRequest, model.Response{
            StatusCode: http.StatusBadRequest,
            Message:    err.Error(),
            Data:       nil,
        })
    }
	req, ok := validatedReq.(*modelinsurance.ReqUserInsurance)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	err = u.CustomDate.UnmarshalJSON([]byte(`"` + req.NgayDong + `"`))
	if err != nil {
		fmt.Println("Error parsing NgayBatDau:", err)
	}
	err = u.CustomDate2.UnmarshalJSON([]byte(`"` + req.NgayHetHan + `"`))
	if err != nil {
		fmt.Println("Error parsing NgayBatDau:", err)
	}
	Insurance := modelinsurance.ResUserInsurance{
		IDNhanVien:  idUser,
		IDBaoHiem:    idInsurance,
		NgayDong:    u.CustomDate.Time ,
		NgayHetHan: u.CustomDate2.Time,
	}
	InsuranceResult, err := u.InsuranceUserRepo.UpdateInsuranceById(c.Request().Context(), Insurance)
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
func (u *InsuranceUserController) DeleteInsuranceUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	idInsurance, err := strconv.Atoi(c.QueryParam("ids"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.InsuranceUserRepo.DeleteInsuranceById(c.Request().Context(), idInsurance,idUser)
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
