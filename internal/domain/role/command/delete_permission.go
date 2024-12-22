package command

// 删除角色权限命令
type DeletePermissionCommand struct {
	Id       string
	PermissionIds []string
}
