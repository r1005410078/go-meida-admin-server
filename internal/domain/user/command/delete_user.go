package command

// 删除用户命令
type DeleteUserCommand struct {
	Id string `json:"id" binding:"required"`
}
