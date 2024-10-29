package middleware

import (
	"BackEnd/mod/banana"
	"BackEnd/mod/db"
	"BackEnd/mod/model"
	modelrole "BackEnd/mod/model/model_role"
	"BackEnd/mod/security"
	"fmt"
	"net/http"
	"reflect"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// kiem tra xem token co hop le khong
func JWTMiddlware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(security.SECRET_KEY),
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			return c.JSON(http.StatusUnauthorized, model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    banana.NotSignIn.Error(),
				Data:       nil,
			})
		},
	}
	return middleware.JWTWithConfig(config)
}

func PermissionMiddlewar(requiredPermission model.Permission) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"status_code": http.StatusUnauthorized,
					"message":     "Missing token",
					"data":        nil,
				})
			}

			if len(token) > 7 && token[:7] == "Bearer " {
				token = token[7:]
			}

			result, _ := security.ExtractClaims(token)
			data := result["Role"].([]interface{})
			var dataRole []int
			for _, role := range data {
				if roleInt, ok := role.(float64); ok {
					dataRole = append(dataRole, int(roleInt))
				}
			}
			var permissions []model.Permission
			for _, role := range dataRole {
				if perms, ok := model.RolePermissions[role]; ok {
					permissions = append(permissions, perms...)
				}
			}

			for _, permission := range permissions {
				if permission == requiredPermission {
					return next(c) // Cho phép tiếp tục nếu quyền hợp lệ
				}
			}
			return c.JSON(http.StatusForbidden, model.Response{
				StatusCode: http.StatusForbidden,
				Message:    banana.TitleErrol.Error(),
				Data:       nil,
			})
		}
	}
}

func PermissionMiddleware(ROLE []int, condition string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sqlDB := db.NewSqlConfig()
			sqlDB.Connect()             
			defer sqlDB.Close()   
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"status_code": http.StatusUnauthorized,
					"message":     "Missing token",
					"data":        nil,
				})
			}

			if len(token) > 7 && token[:7] == "Bearer " {
				token = token[7:]
			}

			result, _ := security.ExtractClaims(token)
			data := result["Role"].([]interface{})
			var dataRole []int
			for _, role := range data {
				if roleInt, ok := role.(float64); ok {
					dataRole = append(dataRole, int(roleInt))
				}
			}

			var roleUser modelrole.ResUserRole // Đảm bảo sử dụng đúng kiểu dữ liệu

			query := `SELECT 
					MAX(Xem) AS Xem,
					MAX(Them) AS Them,
					MAX(Sua) AS Sua,
					MAX(Xoa) AS Xoa
				FROM 
					chucdanh_vaitro
				WHERE 
					IDVaiTro IN (?)
					AND IDChucDanh IN (?)
					`
			query, args, err := sqlx.In(query,ROLE, dataRole)
			if err != nil {
				log.Error(err.Error())
				return c.JSON(http.StatusInternalServerError, model.Response{
					StatusCode: http.StatusInternalServerError,
					Message:    "Failed to build query",
					Data:       nil,
				})
			}
			query = sqlDB.Db.Rebind(query) // Sử dụng sqlDB.Db để tái kết nối
			err = sqlDB.Db.GetContext(c.Request().Context(), &roleUser, query, args...)
			if err != nil {
				log.Error(err.Error())
				return c.JSON(http.StatusInternalServerError, model.Response{
					StatusCode: http.StatusInternalServerError,
					Message:    banana.TitleErrol.Error(),
					Data:       nil,
				})
			}
			fmt.Println(roleUser)
			val := reflect.ValueOf(roleUser)
			field := val.FieldByName(condition)
			fmt.Println(condition)
			fmt.Println(field)
			if field.IsValid() && field.Kind() == reflect.Bool {
				if field.Bool() {
					return next(c)
				}
			}
			return c.JSON(http.StatusForbidden, model.Response{
				StatusCode: http.StatusForbidden,
				Message:    banana.TitleErrol.Error(),
				Data:       nil,
			})
		}
	}
}