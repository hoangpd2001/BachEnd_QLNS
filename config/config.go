package config

import (
	skillcontrollerr "BackEnd/mod/controller/skill_controller"
	typecontroller "BackEnd/mod/controller/type_controller"
	usercontroller "BackEnd/mod/controller/user_controller"
	reposkill "BackEnd/mod/repository/repo_skill"
	repotype "BackEnd/mod/repository/repo_type"
	repoimpl "BackEnd/mod/repository/repo_user/repo_impl"
	"BackEnd/mod/router"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InitApp(e *echo.Echo, sqlDB *sqlx.DB) *router.API {
	userRepo := repoimpl.NewUserRepo(sqlDB)
	educationRepo := repoimpl.NewEducationRepo(sqlDB)
	relativeRepo := repoimpl.NewRelativeRepo(sqlDB)
	typeRepo := repotype.NewTypeRepo(sqlDB)
	skillRepo := reposkill.NewSkillRepo(sqlDB)
	skillUserRepo := reposkill.NewSkillUserRepo(sqlDB)
	userController := usercontroller.UseController{
		UserRepo: userRepo,
	}
	educationController := usercontroller.EducationController{
		EducationRepo: educationRepo,
	}
	relativeController := usercontroller.RelativeController{
		RelativeRepo: relativeRepo,
	}
	typecontroller := typecontroller.TypeController{
		TypeRepo: *typeRepo,
	}
	skillcontroller := skillcontrollerr.SkillController{
		SkillRepo: *skillRepo,
	}
	skillUserController := skillcontrollerr.SkillUserController{
		SkillUserRepo: *skillUserRepo,
	}
	api := &router.API{
		Echo:                e,
		UseController:       userController,
		EducationController: educationController,
		RelativeController:  relativeController,
		TypeController:      typecontroller,
		SkillController:     skillcontroller,
		SkillUserController: skillUserController,
	}

	return api
}
