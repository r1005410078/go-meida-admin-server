package events

import "github.com/r1005410078/meida-admin-server/internal/domain/forms/command"

// 创建表单字段
type CreateFormsFieldsEvent struct {
	*command.SaveFormsFieldsCommand
}

// 创建错误
type CreateFormsFiledsFailedEvent struct {
	*command.SaveFormsFieldsCommand
	Err error
}

// 更新表单字段
type UpdateFormsFieldsEvent struct {
	*command.SaveFormsFieldsCommand
}

// 更新错误
type UpdateFormsFieldsFailedEvent struct {
	*command.SaveFormsFieldsCommand
	Err error
}

type DeleteFormsFiledsEvent struct {
  *command.DeleteFormsFieldsCommand
}

type DeleteFormsFiledsFailedEvent struct {
	*command.DeleteFormsFieldsCommand
	Err error
}