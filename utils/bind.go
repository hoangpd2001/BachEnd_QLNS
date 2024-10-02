
package utils

import (
	"BackEnd/mod/model"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Bind struct{

}
func (b *Bind) BindAndValidate(c echo.Context, req interface{}) (interface{}, error) {
    // Bind dữ liệu từ request
    if err := c.Bind(req); err != nil {
        log.Error(err.Error())
        return nil, c.JSON(http.StatusBadRequest, model.Response{
            StatusCode: http.StatusBadRequest,
            Message:    err.Error(),
            Data:       nil,
        })
    }

    // Validate dữ liệu
    validato := validator.New()
    if err := validato.Struct(req); err != nil {
        log.Error(err.Error())
        return nil, c.JSON(http.StatusBadRequest, model.Response{
            StatusCode: http.StatusBadRequest,
            Message:    err.Error(),
            Data:       nil,
        })
    }

    // Nếu không có lỗi, trả về request data
    return req, nil
}
