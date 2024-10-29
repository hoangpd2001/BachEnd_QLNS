package reqUser

type ReqSignIn struct {
	Email string `validate:"required"`
	Pass  string `validate:"required"`
}
type ReqSignInEdit struct {
	Email string `validate:"required"`
	SDT  string `validate:"required"`
}
