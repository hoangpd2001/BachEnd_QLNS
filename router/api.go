package router

import (
	branchcontroller "BackEnd/mod/controller/branch_controller"
	departmentcontroller "BackEnd/mod/controller/department_controller"
	grupcontroller "BackEnd/mod/controller/grup_controller"
	insurancecontroller "BackEnd/mod/controller/insurance_controller"
	levelcontroller "BackEnd/mod/controller/level_controller"
	rolecontroller "BackEnd/mod/controller/role_controller"
	skillcontroller "BackEnd/mod/controller/skill_controller"
	titlecontroller "BackEnd/mod/controller/title_controller"
	typecontroller "BackEnd/mod/controller/type_controller"
	user_controller "BackEnd/mod/controller/user_controller"
	myMiddleware "BackEnd/mod/middleware"

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
	RoleController          rolecontroller.RoleController
	RoleUserController      rolecontroller.RoleUserController
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-in", api.UseController.HandleSignIn)
	api.Echo.POST("/user/editPass", api.UseController.HanEditLogin)


	user := api.Echo.Group("", myMiddleware.JWTMiddlware())

	user.POST("/user/creatUser", myMiddleware.PermissionMiddleware([]int{2},"Them")(api.UseController.CreatUser))
	user.GET("/user/selectAll", myMiddleware.PermissionMiddleware([]int{2,4,3},"Xem")(api.UseController.SelectUserAll))
	user.GET("/user/selectOne/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.UseController.SelectUserById))
	user.GET("/user/selectCount/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.UseController.SelectCountUser))
	user.PUT("/user/update/", myMiddleware.PermissionMiddleware([]int{2},"Sua")(api.UseController.UpdateUserById))

	user.POST("/user/education/creat", myMiddleware.PermissionMiddleware([]int{2},"Them")(api.EducationController.CreatEducation))
	user.GET("/user/education/SelectAll/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.EducationController.SelectEducationByUser))
	user.GET("/user/education/SelectOne/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.EducationController.SelectEducationById))
	user.PUT("/user/education/Update/", myMiddleware.PermissionMiddleware([]int{2},"Sua")(api.EducationController.UpdateEducationById))
	user.DELETE("/user/education/Delete/", myMiddleware.PermissionMiddleware([]int{2},"Xoa")(api.EducationController.DeleteEducationById))

	user.POST("/user/type/creat", myMiddleware.PermissionMiddleware([]int{2},"Them")(api.TypeController.CreatType))              
	user.GET("/user/type/selectAll", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.TypeController.SelectTypeAll))       
	user.GET("/user/type/selectOne/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.TypeController.SelelectTypeByUser)) 
	user.PUT("/user/type/update/", myMiddleware.PermissionMiddleware([]int{2},"Sua")(api.TypeController.UpdateTypeById))        
	user.DELETE("/user/type/delete/", myMiddleware.PermissionMiddleware([]int{2},"Xoa")(api.TypeController.DeleteTypeById)) 


	user.POST("/user/relative/creat/", myMiddleware.PermissionMiddleware([]int{2},"Them")(api.RelativeController.CreatRelative))
	user.GET("/user/relative/select/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.RelativeController.SelectRelativeByUser))
	user.PUT("/user/relative/update/", myMiddleware.PermissionMiddleware([]int{2},"Sua")(api.RelativeController.UpdateRelativeByUser))

	user.POST("/user/skill/creat", myMiddleware.PermissionMiddleware([]int{4},"Them")(api.SkillController.CreatSkill))
	user.GET("/user/skill/selectAll", myMiddleware.PermissionMiddleware([]int{4},"Xem")(api.SkillController.SelectSkillAll))
	user.GET("/user/skill/selectOne/", myMiddleware.PermissionMiddleware([]int{4},"Xem")(api.SkillController.SelelectSkillById))
	user.PUT("/user/skill/update/", myMiddleware.PermissionMiddleware([]int{4},"Sua")(api.SkillController.UpdateSkillById))
	user.DELETE("/user/skill/delete/", myMiddleware.PermissionMiddleware([]int{4},"Xoa")(api.SkillController.DeleteSkillById))

	user.POST("/user/skilluser/creat/", myMiddleware.PermissionMiddleware([]int{4},"Them")(api.SkillUserController.CreatSkillUser))
	user.GET("/user/skilluser/selectAll/", myMiddleware.PermissionMiddleware([]int{4},"Xem")(api.SkillUserController.SelectSkillUserAll))
	user.GET("/user/skilluser/select/", myMiddleware.PermissionMiddleware([]int{4},"Xem")(api.SkillUserController.SelectSkillUser))
	user.GET("/user/skilluser/selectOne/", myMiddleware.PermissionMiddleware([]int{4},"Xem")(api.SkillUserController.SelelectSkillUser))
	user.PUT("/user/skilluser/update/", myMiddleware.PermissionMiddleware([]int{4},"Sua")(api.SkillUserController.UpdateSkillUser))
	user.DELETE("/user/skilluser/delete/", myMiddleware.PermissionMiddleware([]int{4},"Xoa")(api.SkillUserController.DeleteSkillUser))

	user.POST("/userTitle/creat", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.UserTitleController.CreatUserTitle))
	user.GET("/userTitle/selectAll/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.UserTitleController.SelectUserTitleAll))
	user.GET("/userTitle/selectOne/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.UserTitleController.SelelectUserTitle))
	user.PUT("/userTitle/update/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.UserTitleController.UpdateUserTitle))
	user.DELETE("/userTitle/delete/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.UserTitleController.DeleteUserTitle))

	user.POST("/userGrup/creat", myMiddleware.PermissionMiddleware([]int{3},"Xem")(api.GrupUserController.CreatGrupUser))
	user.GET("/userGrup/selectAll/", myMiddleware.PermissionMiddleware([]int{3},"Xem")(api.GrupUserController.SelectGrupUserAll))
	user.GET("/userGrup/select/", myMiddleware.PermissionMiddleware([]int{3},"Xem")(api.GrupUserController.SelelectGrupUser))

	user.POST("/branch/creat", myMiddleware.PermissionMiddleware([]int{3},"Them")(api.BranchController.CreatBranch))
	user.GET("/branch/selectAll", myMiddleware.PermissionMiddleware([]int{3,2},"Xem")(api.BranchController.SelectBranchAll))
	user.GET("/branch/selectOne/", myMiddleware.PermissionMiddleware([]int{3},"Xem")(api.BranchController.SelelectBranchById))
	user.PUT("/branch/update/", myMiddleware.PermissionMiddleware([]int{3},"Sua")(api.BranchController.UpdateBranchById))
	user.DELETE("/branch/delete/", myMiddleware.PermissionMiddleware([]int{3},"Xoa")(api.BranchController.DeleteBranchById))

	user.POST("/title/creat", myMiddleware.PermissionMiddleware([]int{2},"Them")(api.TitleController.CreatTitle))
	user.GET("/title/selectAll", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.TitleController.SelectTitleAll))
	user.GET("/title/selectOne/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.TitleController.SelelectTitleById))
	user.PUT("/title/update/", myMiddleware.PermissionMiddleware([]int{2},"Sua")(api.TitleController.UpdateTitleById))
	user.DELETE("/title/delete/", myMiddleware.PermissionMiddleware([]int{2},"Xoa")(api.TitleController.DeleteTitleById))

	user.POST("/department/creat", myMiddleware.PermissionMiddleware([]int{3},"Them")(api.DepartmentController.CreatDepartment))
	user.GET("/department/selectAll", myMiddleware.PermissionMiddleware([]int{3,2},"Xem")(api.DepartmentController.SelectDepartmentAll))
	user.GET("/department/selectOne/", myMiddleware.PermissionMiddleware([]int{3,2},"Xem")(api.DepartmentController.SelelectDepartmentById))
	user.GET("/department/selectByBranch/", myMiddleware.PermissionMiddleware([]int{3,2},"Xem")(api.DepartmentController.SelelectDepartmentByBranch))
	user.PUT("/department/update/", myMiddleware.PermissionMiddleware([]int{3},"Sua")(api.DepartmentController.UpdateDepartmentById))
	user.DELETE("/department/delete/", myMiddleware.PermissionMiddleware([]int{3},"Xoa")(api.DepartmentController.DeleteDepartment))

	user.POST("/level/creat", myMiddleware.PermissionMiddleware([]int{3},"Them")(api.LevelController.CreatLevel))
	user.GET("/level/selectAll", myMiddleware.PermissionMiddleware([]int{3,2},"Xem")(api.LevelController.SelectLevelAll))
	user.GET("/level/selectOne/", myMiddleware.PermissionMiddleware([]int{3,2},"Xem")(api.LevelController.SelelectLevelByUser))
	user.PUT("/level/update/", myMiddleware.PermissionMiddleware([]int{3},"Sua")(api.LevelController.UpdateLevelById))
	user.DELETE("/level/delete/", myMiddleware.PermissionMiddleware([]int{3},"Xoa")(api.LevelController.DeleteLevelById))

	user.POST("/grup/creat", myMiddleware.PermissionMiddleware([]int{3},"Them")(api.GrupController.CreatGrup))
	user.GET("/grup/selectAll", myMiddleware.PermissionMiddleware([]int{3},"Xem")(api.GrupController.SelectGrupAll))
	user.PUT("/grup/update/", myMiddleware.PermissionMiddleware([]int{3},"Xem")(api.GrupController.UpdateGrupById))
	user.DELETE("/grup/delete/", myMiddleware.PermissionMiddleware([]int{3},"Xoa")(api.GrupController.DeleteGrupById))
	user.DELETE("/userGrup/delete/", myMiddleware.PermissionMiddleware([]int{3},"Xoa")(api.GrupUserController.DeleteGrupUser))

	user.POST("/user/insurance/creat", myMiddleware.PermissionMiddleware([]int{2},"Them")(api.InsuranceController.CreatInsurance))
	user.GET("/user/insurance/selectAll", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.InsuranceController.SelectInsuranceAll))
	user.GET("/user/insurance/selectOne/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.InsuranceController.SelelectInsuranceById))
	user.PUT("/user/insurance/update/", myMiddleware.PermissionMiddleware([]int{2},"Sua")(api.InsuranceController.UpdateInsuranceById))
	user.DELETE("/user/insurance/delete/", myMiddleware.PermissionMiddleware([]int{2},"Xoa")(api.InsuranceController.DeleteInsuranceById))

	user.POST("/user/insuranceUser/creat/", myMiddleware.PermissionMiddleware([]int{2},"Them")(api.InsuranceUserController.CreatInsuranceUser))
	user.GET("/user/insuranceUser/selectAll/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.InsuranceUserController.SelectInsuranceUserAll))
	user.GET("/user/insuranceUser/select/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.InsuranceUserController.SelectInsuranceUser))
	user.GET("/user/insuranceUser/selectOne/", myMiddleware.PermissionMiddleware([]int{2},"Xem")(api.InsuranceUserController.SelelectInsuranceUserOne))
	user.PUT("/user/insuranceUser/update/", myMiddleware.PermissionMiddleware([]int{2},"Sua")(api.InsuranceUserController.UpdateInsuranceUser))
	user.DELETE("/user/insuranceUser/delete/", myMiddleware.PermissionMiddleware([]int{2},"Xoa")(api.InsuranceUserController.DeleteInsuranceUser))

	user.GET("/role/selectAll", myMiddleware.PermissionMiddleware([]int{1},"Xem")(api.RoleController.SelectRoleAll))

	user.POST("/userrole/creat/", myMiddleware.PermissionMiddleware([]int{1},"Them")(api.RoleUserController.CreatRoleUser))
	user.GET("/userrole/selectAll", myMiddleware.PermissionMiddleware([]int{1},"Xem")(api.RoleUserController.SelectRoleUserAll))
	// user.GET("/userrole/selectAll", myMiddleware.PermissionMiddleware([]int{1},"Xem")(api.RoleUserController.SelelectRoleUser())
	user.PUT("/userrole/update/", myMiddleware.PermissionMiddleware([]int{1},"Sua")(api.RoleUserController.UpdateRoleUser))
	user.DELETE("/userrole/delete/", myMiddleware.PermissionMiddleware([]int{1},"Xoa")(api.RoleUserController.DeleteRoleUser))

}
