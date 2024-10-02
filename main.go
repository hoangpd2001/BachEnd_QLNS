package main

import (
	"BackEnd/mod/config"
	"BackEnd/mod/db"

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

	api := config.InitApp(e, sql.Db)
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1323"))
}
