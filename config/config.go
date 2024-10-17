package config

import (
	branchcontroller "BackEnd/mod/controller/branch_controller"
	departmentcontroller "BackEnd/mod/controller/department_controller"
	grupcontroller "BackEnd/mod/controller/grup_controller"
	insurancecontroller "BackEnd/mod/controller/insurance_controller"
	levelcontroller "BackEnd/mod/controller/level_controller"
	skillcontrollerr "BackEnd/mod/controller/skill_controller"
	titlecontroller "BackEnd/mod/controller/title_controller"
	typecontroller "BackEnd/mod/controller/type_controller"
	usercontroller "BackEnd/mod/controller/user_controller"
	repoinsurance "BackEnd/mod/repository/repo_insurance"
	repobranch "BackEnd/mod/repository/repo_branch"
	repodepartment "BackEnd/mod/repository/repo_department"
	repogrup "BackEnd/mod/repository/repo_grup"
	repolevel "BackEnd/mod/repository/repo_level"
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
	levelRepo := repolevel.NewLevelRepo(sqlDB)
	grupRepo := repogrup.NewGrupRepo(sqlDB)
	userGrupRepo := repogrup.NewGrupUserRepo(sqlDB)
	insuranceRepo := repoinsurance.NewInsuranceRepo(sqlDB)
	insuranceUserRepo := repoinsurance.NewInsuranceUserRepo(sqlDB)

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
	levelcontroller := levelcontroller.LevelController{
		LevelRepo: *levelRepo,
	}
	userGrupcontroller := grupcontroller.GrupUserController{
		GrupUserRepo: *userGrupRepo,
	}
	grupcontroller := grupcontroller.GrupController{
		GrupRepo: *grupRepo,
	}
	insuranceUserController := insurancecontroller.InsuranceUserController{
		InsuranceUserRepo: *insuranceUserRepo,
	}
	insurancecontroller := insurancecontroller.InsuranceController{
		InsuranceRepo: *insuranceRepo,
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
		LevelController:      levelcontroller,
		GrupController:       grupcontroller,
		GrupUserController:   userGrupcontroller,
		InsuranceController:  insurancecontroller,
		InsuranceUserController: insuranceUserController,
	}

	return api
}
