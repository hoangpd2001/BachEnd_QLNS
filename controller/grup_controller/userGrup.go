package grupcontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	modelgrup "BackEnd/mod/model/model_grup"
	repogrup "BackEnd/mod/repository/repo_grup"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GrupUserController struct {
	GrupUserRepo repogrup.GrupUserRepo
	Bind         utils.Bind
	CustomDate   utils.CustomDate
}

// ========================================================================================================
func (u *GrupUserController) CreatGrupUser(c echo.Context) error {
	req := &modelgrup.ReqUserGrup{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req, ok := validatedReq.(*modelgrup.ReqUserGrup)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	res := modelgrup.ResUserGrup{
		IDNhanVien:  req.IDNhanVien,
		IDNhom: req.IDNhom,
	}
	skillResult, err := u.GrupUserRepo.CreatGrupUser(c.Request().Context(), res)
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

func (u *GrupUserController) SelectGrupUserAll(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	skillResult, err := u.GrupUserRepo.SelectGrupUserByGrup(c.Request().Context(), idUser)
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

func (u *GrupUserController) SelelectGrupUser(c echo.Context) error {
	

	skillResult, err := u.GrupUserRepo.SelelectGrupUseeAll(c.Request().Context())
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

// func (u *GrupUserController) UpdateGrupUser(c echo.Context) error {
// 	idUser, err := strconv.Atoi(c.QueryParam("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusConflict, model.Response{
// 			StatusCode: http.StatusConflict,
// 			Message:    banana.GetIdFailed.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	idGrup, err := strconv.Atoi(c.QueryParam("ids"))
// 	if err != nil {
// 		return c.JSON(http.StatusConflict, model.Response{
// 			StatusCode: http.StatusConflict,
// 			Message:    banana.GetIdFailed.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	req := &modelgrup.ReqUserGrup{}
// 	validatedReq, err := u.Bind.BindAndValidate(c, req)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, model.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}
// 	req, ok := validatedReq.(*modelgrup.ReqUserGrup)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, model.Response{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    "Failed to cast validated request",
// 			Data:       nil,
// 		})
// 	}
// 	err = u.CustomDate.UnmarshalJSON([]byte(req.NgayDanhGia))
// 	skill := modelgrup.ResUserGrup{
// 		IDNhanVien:  idUser,
// 		IDKyNang:    idGrup,
// 		MucDo:       req.MucDo,
// 		NgayDanhGia: u.CustomDate.Time,
// 		IDKyNangMoi: req.IDKyNang,
// 	}
// 	skillResult, err := u.GrupUserRepo.UpdateSkillById(c.Request().Context(), skill)
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

// ====================================================================================================================
func (u *GrupUserController) DeleteGrupUser(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	idGrup, err := strconv.Atoi(c.QueryParam("ids"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}

	result, err := u.GrupUserRepo.DeleteGrupUser(c.Request().Context(), idGrup, idUser)
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
