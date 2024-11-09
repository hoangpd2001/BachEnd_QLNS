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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	api := config.InitApp(e, sql.Db)
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1323"))
	//e.Logger.Fatal(e.Start("0.0.0.0:" + os.Getenv("PORT")))
}
