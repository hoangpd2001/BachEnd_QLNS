package config

import (
	usercontroller "BackEnd/mod/controller/user_controller"
	repoimpl "BackEnd/mod/repository/repo_user/repo_impl"
	"BackEnd/mod/router"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InitApp(e *echo.Echo, sqlDB *sqlx.DB) *router.API {
	userRepo := repoimpl.NewUserRepo(sqlDB)
	educationRepo := repoimpl.NewEducationRepo(sqlDB)
	relativeRepo := repoimpl.NewRelativeRepo(sqlDB)
	userController := usercontroller.UseController{
		UserRepo: userRepo,
	}
	educationController := usercontroller.EducationController{
		EducationRepo: educationRepo,
	}
	relativeController := usercontroller.RelativeController{
		RelativeRepo: relativeRepo,
	}
	api := &router.API{
		Echo:                e,
		UseController:       userController,
		EducationController: educationController,
		RelativeController:  relativeController,
	}

	return api
}
