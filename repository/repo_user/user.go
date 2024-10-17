package repouser

import (
	resUser "BackEnd/mod/model/model_user/res_user"
	"context"
	"database/sql"
)

type UserRepo interface {
	CreatUser(context context.Context, user resUser.ResUser) (resUser.ResUser, error)
	// CheckLogin(context context.Context, loginReq req_user.ReqSignIn) (resUser.User, error)
	SelectUserAll(context context.Context) ([]resUser.ResUser, error)
	SelectUserById(context context.Context, IdUser int) (resUser.ResUser, error)
	// UpdateUserById(context context.Context, user model.User ) (model.User, error)
}
type EducationRepo interface {
	CreatEducation(context context.Context, UserEdu resUser.ResEducation) (resUser.ResEducation, error)
	SelectEducationByUser(context context.Context, UserId int) ([]resUser.ResEducation, error)
	SelectEducationById(context context.Context, EducationId int) (resUser.ResEducation, error)
	UpdateEducationById(context context.Context, Education resUser.ResEducation) (resUser.ResEducation, error)
	DeleteEducationById(context context.Context, EducationId int) (sql.Result, error)
}
type RelativeRepo interface {
	CreatRelative(context context.Context, UserRelative resUser.ResRelative) (resUser.ResRelative, error)
	SelectRelativeByUser(context context.Context, UserId int) ([]resUser.ResRelative, error)
	UpdateRelativeByUser(context context.Context, Relative resUser.ResRelative) (resUser.ResRelative, error)
}