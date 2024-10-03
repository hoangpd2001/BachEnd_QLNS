package main

import (
	"BackEnd/mod/config"
	"BackEnd/mod/db"
	"net/http"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	sql := db.NewSqlConfig()
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*") // Cho phép từ mọi miền
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

			// Kiểm tra nếu là phương thức OPTIONS
			if c.Request().Method == http.MethodOptions {
				return c.NoContent(http.StatusNoContent)
			}
			return next(c)
		}
	}))

	api := config.InitApp(e, sql.Db)
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1323"))
}
