package banana

import "errors"

var (
	UserConflict = errors.New("Người dùng đã tồn tại ")
	SignUpFail   = errors.New("Đăng kí thất bại")
	SignInFail   = errors.New("Tài khoản hoặc mật khẩu không chính xác")
	EditLoginFail   = errors.New("SDT hoặc Email không chính xác")
	UserNotFound = errors.New("Người dùng không tồn tại")
	NotSignIn    = errors.New("Vui lòng đăng nhập")
	TitleErrol	=errors.New("Vui lòng đăng nhập với phân quyền cao hơn")
	UpdateFailed = errors.New("Cập nhật thông tin thất bại ")
	GetIdFailed  = errors.New("Lấy thông tin thất bại")
	ForenkeyErrol = errors.New("Dữ liệu này đang được liên kết với dữ liệu khác")
	SameName = errors.New("Dữ liệu đã tồn tại")
	SererError = errors.New("Server đang gặp lỗi, vui lòng thử lại")

)
