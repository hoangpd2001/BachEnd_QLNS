package banana

import "errors"

var (
	UserConflict = errors.New("Người dùng đã tồn tại ")
	SignUpFail   = errors.New("Đăng kí thất bại")
	UserNotFound = errors.New("Người dùng không tồn tại")
	NotSignIn    = errors.New("Vui lòng đăng nhập")
	UpdateFailed = errors.New("Cập nhật thông tin thất bại ")
	GetIdFailed  = errors.New("Lấy thông tin người dùng thất bại")
)
