package resuser

type ResSingin struct {
	ID             int       `db:"ID"             json:"ID,omitempty"`
	Email string `db:"Email"  json:"Email,omitempty"`
	Role         int `db:"IDChucDanh"           json:"Role,"`
	MatKhau         string `db:"MatKhau"           json:"MatKhau,omitempty"`
	Token string `                    json:"Token,omitempty"`
}
type ResSignInEdit struct {
	Email			string `db:"Email"  json:"Email,omitempty"`
	MatKhau         string `db:"MatKhau"           json:"MatKhau,omitempty"`
}