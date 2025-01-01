package events

import "github.com/r1005410078/meida-admin-server/internal/domain/forms/command"

type CreateFormsEvent struct {
	*command.SaveFormsCommand
}

type UpdateFormsEvent struct {
	*command.SaveFormsCommand
}

// 保存错误
type SaveFormsFailedEvent struct {
	*command.SaveFormsCommand
	Err error
}

type DeleteFormsCommand struct {
	Id *string
}

// 删除错误
type DeleteFormsFailedEvent struct {
	*command.DeleteFormsCommand
	Err error
}

