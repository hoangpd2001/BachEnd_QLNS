package middleware

import "github.com/labstack/echo/v4"

func IsAdmin() echo.MiddlewareFunc{
	return func(next echo.HandlerFunc) echo.HandlerFunc{
		return func(c echo.Context) error{
			if 1==1{
				return c.JSON(500, echo.Map{
					"err":"ban chua dang nhap",
				})
			}		
			return next(c)
		}
	}
}