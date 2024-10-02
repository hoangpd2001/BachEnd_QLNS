package router

import (
	user_controller "BackEnd/mod/controller/user_controller"
	// myMiddleware "BackEnd/mod/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo                *echo.Echo
	UseController       user_controller.UseController
	EducationController user_controller.EducationController
	RelativeController  user_controller.RelativeController
}

func (api *API) SetupRouter() {
	// api.Echo.POST("/user/sign-in", api.UseController.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UseController.CreatUser)
	api.Echo.GET("/user/selectAll", api.UseController.SelectUserAll)
	api.Echo.GET("/user/selectOne/", api.UseController.SelectUserById)
//	api.Echo.PUT("/profile/update", api.UseController.UpdateUserById)
	// user := api.Echo.Group("/user", myMiddleware.JWTMiddlware())
	// user.GET("/profile", api.UseController.Profile)
	// user.PUT("/profile/update", api.UseController.UpdateProfile)

	api.Echo.POST("/user/education/creat", api.EducationController.CreatEducation)
	api.Echo.GET("/user/education/SelectAll/", api.EducationController.SelectEducationByUser)
	api.Echo.GET("/user/education/SelectOne/", api.EducationController.SelectEducationById)
	api.Echo.PUT("/user/education/Update", api.EducationController.UpdateEducationById)

	
	api.Echo.POST("/user/relative/creat/", api.RelativeController.CreatRelative)
	api.Echo.GET("/user/relative/select/", api.RelativeController.SelectRelativeByUser)
	api.Echo.PUT("/user/relative/update/", api.RelativeController.UpdateRelativeByUser)

}
