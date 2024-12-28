package command

import "github.com/r1005410078/meida-admin-server/internal/domain/user/events"

type UserStatusCommand struct {
	Id string `json:"id" binding:"required"`
	// 用户状态
	Status *string `json:"status" binding:"required"`
}

func (command *UserStatusCommand) ToEvent() *events.UserStatusEvent {
	return &events.UserStatusEvent{
		Id: command.Id,
		Status: command.Status,
	}
}