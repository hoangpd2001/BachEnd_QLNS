package router

import (
	branchcontroller "BackEnd/mod/controller/branch_controller"
	departmentcontroller "BackEnd/mod/controller/department_controller"
	grupcontroller "BackEnd/mod/controller/grup_controller"
	insurancecontroller "BackEnd/mod/controller/insurance_controller"
	levelcontroller "BackEnd/mod/controller/level_controller"
	skillcontroller "BackEnd/mod/controller/skill_controller"
	titlecontroller "BackEnd/mod/controller/title_controller"
	typecontroller "BackEnd/mod/controller/type_controller"
	user_controller "BackEnd/mod/controller/user_controller"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo                    *echo.Echo
	UseController           user_controller.UseController
	EducationController     user_controller.EducationController
	RelativeController      user_controller.RelativeController
	TypeController          typecontroller.TypeController
	SkillController         skillcontroller.SkillController
	SkillUserController     skillcontroller.SkillUserController
	BranchController        branchcontroller.BranchController
	TitleController         titlecontroller.TitleController
	DepartmentController    departmentcontroller.DepartmentController
	UserTitleController     titlecontroller.UserTitleController
	LevelController         levelcontroller.LevelController
	GrupController          grupcontroller.GrupController
	GrupUserController      grupcontroller.GrupUserController
	InsuranceController     insurancecontroller.InsuranceController
	InsuranceUserController insurancecontroller.InsuranceUserController
}

func (api *API) SetupRouter() {
	// api.Echo.POST("/user/sign-in", api.UseController.HandleSignIn)
	api.Echo.POST("/user/creatUser", api.UseController.CreatUser)
	api.Echo.GET("/user/selectAll", api.UseController.SelectUserAll)
	api.Echo.GET("/user/selectOne/", api.UseController.SelectUserById)
	//	api.Echo.PUT("/profile/update", api.UseController.UpdateUserById)
	// user := api.Echo.Group("/user", myMiddleware.JWTMiddlware())
	// user.GET("/profile", api.UseController.Profile)
	// user.PUT("/profile/update", api.UseController.UpdateProfile)

	api.Echo.POST("/user/education/creat", api.EducationController.CreatEducation)
	api.Echo.GET("/user/education/SelectAll/", api.EducationController.SelectEducationByUser)
	api.Echo.GET("/user/education/SelectOne/", api.EducationController.SelectEducationById)
	api.Echo.PUT("/user/education/Update/", api.EducationController.UpdateEducationById)
	api.Echo.DELETE("/user/education/Delete/", api.EducationController.DeleteEducationById)

	api.Echo.POST("/user/relative/creat/", api.RelativeController.CreatRelative)
	api.Echo.GET("/user/relative/select/", api.RelativeController.SelectRelativeByUser)
	api.Echo.PUT("/user/relative/update/", api.RelativeController.UpdateRelativeByUser)

	api.Echo.POST("/user/type/creat", api.TypeController.CreatType)
	api.Echo.GET("/user/type/selectAll", api.TypeController.SelectTypeAll)
	api.Echo.GET("/user/type/selectOne/", api.TypeController.SelelectTypeByUser)
	api.Echo.PUT("/user/type/update/", api.TypeController.UpdateTypeById)
	api.Echo.DELETE("/user/type/delete/", api.TypeController.DeleteTypeById)

	api.Echo.POST("/user/skill/creat", api.SkillController.CreatSkill)
	api.Echo.GET("/user/skill/selectAll", api.SkillController.SelectSkillAll)
	api.Echo.GET("/user/skill/selectOne/", api.SkillController.SelelectSkillById)
	api.Echo.PUT("/user/skill/update/", api.SkillController.UpdateSkillById)
	api.Echo.DELETE("/user/skill/delete/", api.SkillController.DeleteSkillById)

	api.Echo.POST("/user/skilluser/creat/", api.SkillUserController.CreatSkillUser)
	api.Echo.GET("/user/skilluser/selectAll/", api.SkillUserController.SelectSkillUserAll)
	api.Echo.GET("/user/skilluser/select/", api.SkillUserController.SelectSkillUser)
	api.Echo.GET("/user/skilluser/selectOne/", api.SkillUserController.SelelectSkillUser)
	api.Echo.PUT("/user/skilluser/update/", api.SkillUserController.UpdateSkillUser)
	api.Echo.DELETE("/user/skilluser/delete/", api.SkillUserController.DeleteSkillUser)

	api.Echo.POST("/branch/creat", api.BranchController.CreatBranch)
	api.Echo.GET("/branch/selectAll", api.BranchController.SelectBranchAll)
	api.Echo.GET("/branch/selectOne/", api.BranchController.SelelectBranchById)
	api.Echo.PUT("/branch/update/", api.BranchController.UpdateBranchById)
	api.Echo.DELETE("/branch/delete/", api.BranchController.DeleteBranchById)

	api.Echo.POST("/title/creat", api.TitleController.CreatTitle)
	api.Echo.GET("/title/selectAll", api.TitleController.SelectTitleAll)
	api.Echo.GET("/title/selectOne/", api.TitleController.SelelectTitleById)
	api.Echo.PUT("/title/update/", api.TitleController.UpdateTitleById)
	api.Echo.DELETE("/title/delete/", api.TitleController.DeleteTitleById)

	api.Echo.POST("/department/creat", api.DepartmentController.CreatDepartment)
	api.Echo.GET("/department/selectAll", api.DepartmentController.SelectDepartmentAll)
	api.Echo.GET("/department/selectOne/", api.DepartmentController.SelelectDepartmentById)
	api.Echo.GET("/department/selectByBranch/", api.DepartmentController.SelelectDepartmentByBranch)
	api.Echo.PUT("/department/update/", api.DepartmentController.UpdateDepartmentById)
	api.Echo.DELETE("/department/delete/", api.DepartmentController.DeleteDepartment)

	api.Echo.POST("/userTitle/creat", api.UserTitleController.CreatUserTitle)
	api.Echo.GET("/userTitle/selectAll/", api.UserTitleController.SelectUserTitleAll)
	api.Echo.GET("/userTitle/selectOne/", api.UserTitleController.SelelectUserTitle)
	api.Echo.PUT("/userTitle/update/", api.UserTitleController.UpdateUserTitle)
	api.Echo.DELETE("/userTitle/delete/", api.UserTitleController.DeleteUserTitle)

	api.Echo.POST("/level/creat", api.LevelController.CreatLevel)
	api.Echo.GET("/level/selectAll", api.LevelController.SelectLevelAll)
	api.Echo.GET("/level/selectOne/", api.LevelController.SelelectLevelByUser)
	api.Echo.PUT("/level/update/", api.LevelController.UpdateLevelById)
	api.Echo.DELETE("/level/delete/", api.LevelController.DeleteLevelById)

	api.Echo.POST("/grup/creat", api.GrupController.CreatGrup)
	api.Echo.GET("/grup/selectAll", api.GrupController.SelectGrupAll)
	api.Echo.PUT("/grup/update/", api.GrupController.UpdateGrupById)
	api.Echo.DELETE("/grup/delete/", api.GrupController.DeleteGrupById)

	api.Echo.POST("/userGrup/creat", api.GrupUserController.CreatGrupUser)
	api.Echo.GET("/userGrup/selectAll/", api.GrupUserController.SelectGrupUserAll)
	api.Echo.GET("/userGrup/select/", api.GrupUserController.SelelectGrupUser)
	//	api.Echo.PUT("/userGrup/update/", api.GrupUserController.)
	api.Echo.DELETE("/userGrup/delete/", api.GrupUserController.DeleteGrupUser)

	api.Echo.POST("/user/insurance/creat", api.InsuranceController.CreatInsurance)
	api.Echo.GET("/user/insurance/selectAll", api.InsuranceController.SelectInsuranceAll)
	api.Echo.GET("/user/insurance/selectOne/", api.InsuranceController.SelelectInsuranceById)
	api.Echo.PUT("/user/insurance/update/", api.InsuranceController.UpdateInsuranceById)
	api.Echo.DELETE("/user/insurance/delete/", api.InsuranceController.DeleteInsuranceById)

	api.Echo.POST("/user/insuranceUser/creat/", api.InsuranceUserController.CreatInsuranceUser)
	api.Echo.GET("/user/insuranceUser/selectAll/", api.InsuranceUserController.SelectInsuranceUserAll)
	api.Echo.GET("/user/insuranceUser/select/", api.InsuranceUserController.SelectInsuranceUser)
	api.Echo.GET("/user/insuranceUser/selectOne/", api.InsuranceUserController.SelelectInsuranceUserOne)
	api.Echo.PUT("/user/insuranceUser/update/", api.InsuranceUserController.UpdateInsuranceUser)
	api.Echo.DELETE("/user/insuranceUser/delete/", api.InsuranceUserController.DeleteInsuranceUser)

}
