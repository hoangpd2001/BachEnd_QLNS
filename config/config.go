package config

import (
	branchcontroller "BackEnd/mod/controller/branch_controller"
	departmentcontroller "BackEnd/mod/controller/department_controller"
	skillcontrollerr "BackEnd/mod/controller/skill_controller"
	titlecontroller "BackEnd/mod/controller/title_controller"
	typecontroller "BackEnd/mod/controller/type_controller"
	usercontroller "BackEnd/mod/controller/user_controller"
	repobranch "BackEnd/mod/repository/repo_branch"
	repodepartment "BackEnd/mod/repository/repo_department"
	reposkill "BackEnd/mod/repository/repo_skill"
	repotitle "BackEnd/mod/repository/repo_title"
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
	branchRepo := repobranch.NewBranchRepo(sqlDB)
	titleRepo := repotitle.NewTitleRepo(sqlDB)
	departmentRepo := repodepartment.NewDepartment(sqlDB)
	userTitleRepo := repotitle.NewUserTitleRepo(sqlDB)

	
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
	branchcontroller := branchcontroller.BranchController{
		BranchRepo: *branchRepo,
	}
	UserTitleController := titlecontroller.UserTitleController{
		UserTitleRepo: *userTitleRepo,
	}
	titlecontroller := titlecontroller.TitleController{
		TitleRepo: *titleRepo,
	}
	departmentcontroller := departmentcontroller.DepartmentController{
		DepartmentRepo: *departmentRepo,
	}

	api := &router.API{
		Echo:                 e,
		UseController:        userController,
		EducationController:  educationController,
		RelativeController:   relativeController,
		TypeController:       typecontroller,
		SkillController:      skillcontroller,
		SkillUserController:  skillUserController,
		BranchController:     branchcontroller,
		TitleController:      titlecontroller,
		DepartmentController: departmentcontroller,
		UserTitleController:  UserTitleController,
	}

	return api
}
