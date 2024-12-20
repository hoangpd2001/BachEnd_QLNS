package repouser

import (
	reqUser "BackEnd/mod/model/model_user/req_user"
	resUser "BackEnd/mod/model/model_user/res_user"
	"context"
	"database/sql"
)

type UserRepo interface {
	CreatUser(context context.Context, user resUser.ResUser) (resUser.ResUser, error)
	CheckLogin(context context.Context, loginReq reqUser.ReqSignIn) ([]resUser.ResSingin, error)
	EditLogin(context context.Context, loginReq reqUser.ReqSignInEdit,mk string) (sql.Result, error)
	SelectUserAll(context context.Context) ([]resUser.ResUser, error)
	SelectUserById(context context.Context, IdUser int) (resUser.ResUser, error)
	SelectCountUser(context context.Context) ([]resUser.ResUserCount, error)
	 UpdateUserById(context context.Context, user resUser.ResUser ) (resUser.ResUser, error)
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