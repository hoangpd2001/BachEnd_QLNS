package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Bind struct {
}

func (b *Bind) BindAndValidate(c echo.Context, req interface{}) (interface{}, error) {
	// Bind dữ liệu từ request
	if err := c.Bind(req); err != nil {
		log.Error(err.Error())
		return nil, fmt.Errorf("Dữ liệu không chính xác")
	}

	// Validate dữ liệu
	validato := validator.New()
	if err := validato.Struct(req); err != nil {
		log.Error(err.Error())
		return nil, fmt.Errorf("Dữ liệu không hợp lệ")
	}

	// Nếu không có lỗi, trả về request data
	return req, nil
}
