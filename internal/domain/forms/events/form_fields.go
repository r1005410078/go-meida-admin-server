package events

import "github.com/r1005410078/meida-admin-server/internal/domain/forms/command"

// 创建表单字段
type CreateFormsFiledsEvent struct {
	*command.SaveFormsFiledsCommand
}

// 创建错误
type CreateFormsFiledsFailedEvent struct {
	*command.SaveFormsFiledsCommand
	Err error
}

// 更新表单字段
type UpdateFormsFiledsEvent struct {
	*command.SaveFormsFiledsCommand
}

// 更新错误
type UpdateFormsFiledsFailedEvent struct {
	*command.SaveFormsFiledsCommand
	Err error
}

type DeleteFormsFiledsEvent struct {
	Id string `json:"id" binding:"required"`
}