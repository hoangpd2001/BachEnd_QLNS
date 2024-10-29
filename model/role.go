package model


type Permission string

const (
    View Permission = "VIEW Nhan Su"
    Edit Permission = "EDIT"
    Delete Permission = "DELETE"
    // Thêm các quyền khác tùy theo yêu cầu
)

var RolePermissions = map[int][]Permission{
    5:  {View, Edit, Delete},
    3: {View},

}