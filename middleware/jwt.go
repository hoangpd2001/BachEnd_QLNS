package middleware

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/model"
	"BackEnd/mod/security"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
// kiem tra xem token co hop le khong
func JWTMiddlware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims: &model.JwtCustomClaims{},
		SigningKey: []byte(security.SECRET_KEY),
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": banana.NotSignIn.Error(),
			})
		},
	}
	return middleware.JWTWithConfig(config)
}
