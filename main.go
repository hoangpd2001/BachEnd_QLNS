package main

import (
	"BackEnd/mod/config"
	"BackEnd/mod/db"
	"net/http"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
//	"os"
)

	var E *echo.Echo
func main() {

	sql := db.NewSqlConfig()
	sql.Connect()
	defer sql.Close()
	E = echo.New() 
	E.Use(middleware.AddTrailingSlash())
	E.Use(middleware.Logger())
	E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	api := config.InitApp(E, sql.Db)
	api.SetupRouter()

	E.Logger.Fatal(E.Start(":1323"))
	//e.Logger.Fatal(e.Start("0.0.0.0:" + os.Getenv("PORT")))
}
func GetEcho() *echo.Echo{
	return E
}