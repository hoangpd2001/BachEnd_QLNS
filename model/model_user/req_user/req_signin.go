package reqUser

type ReqSignIn struct {
	Email string `validate:"required"`
	Pass  string `validate:"required"`
}
