package titlecontroller

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	modetitle "BackEnd/mod/model/model_title"
	repotitle "BackEnd/mod/repository/repo_title"
	"BackEnd/mod/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserTitleController struct {
	UserTitleRepo repotitle.UserTitleRepo
	Bind          utils.Bind
	CustomDate    utils.CustomDate
	CustomDate2    utils.CustomDate
}

// ========================================================================================================
func (u *UserTitleController) CreatUserTitle(c echo.Context) error {

	req := &modetitle.ReqUserTitle{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modetitle.ReqUserTitle)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
		err = u.CustomDate.UnmarshalJSON([]byte(req.NgayBatDau))
	if req.NgayKetThuc != ""{
		err = u.CustomDate2.UnmarshalJSON([]byte(req.NgayKetThuc))
	}
	res := modetitle.ResUserTitle{
		IDNhanVien: req.IDNhanVien,
		IDChucDanh: req.IDChucDanh,
		IDPhongBan: req.IDPhongBan,
		NgayBatDau: u.CustomDate.Time,
		NgayKetThuc: u.CustomDate2.Time,
	}
	Result, err := u.UserTitleRepo.CreatUserTitle(c.Request().Context(), res)
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

func (u *UserTitleController) SelectUserTitleAll(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    banana.GetIdFailed.Error(),
			Data:       nil,
		})
	}
	Result, err := u.UserTitleRepo.SelectUserTitleAll(c.Request().Context(), idUser)
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

func (u *UserTitleController) SelelectUserTitle(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
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

	Result, err := u.UserTitleRepo.SelelectTitleByTitle(c.Request().Context(), idTitle, idUser)
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

//========================================================================================================

func (u *UserTitleController) UpdateUserTitle(c echo.Context) error {
	
	req := &modetitle.ReqUserTitle{}
	validatedReq, err := u.Bind.BindAndValidate(c, req)
	if err != nil {
		return err
	}
	req, ok := validatedReq.(*modetitle.ReqUserTitle)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to cast validated request",
			Data:       nil,
		})
	}
	err = u.CustomDate.UnmarshalJSON([]byte(req.NgayBatDau))
	if req.NgayKetThuc != ""{
		err = u.CustomDate2.UnmarshalJSON([]byte(req.NgayKetThuc))
	}
	res := modetitle.ResUserTitle{
		IDNhanVien: req.IDNhanVien,
		IDChucDanh: req.IDChucDanh,
		IDPhongBan: req.IDPhongBan,
		NgayBatDau: u.CustomDate.Time,
		NgayKetThuc: u.CustomDate2.Time,
	}
	Result, err := u.UserTitleRepo.UpdateTitleById(c.Request().Context(), res)
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

// ====================================================================================================================
func (u *UserTitleController) DeleteUserTitle(c echo.Context) error {
	idUser, err := strconv.Atoi(c.QueryParam("id"))
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

	result, err := u.UserTitleRepo.DeleteTitleById(c.Request().Context(), idTitle, idUser)
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
