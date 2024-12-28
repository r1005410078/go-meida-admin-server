package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
)

type DeleteUserCommandHandler struct {
	repo user.IUserAggregateRepository
	eventBus shared.IEventBus
}

func NewDeleteUserCommandHandler(repo user.IUserAggregateRepository, eventBus shared.IEventBus) *DeleteUserCommandHandler {
	return &DeleteUserCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *DeleteUserCommandHandler) Handle(command *command.DeleteUserCommand) error {
	// 如果不是是管理员，不允许删除
	if !h.repo.IsAdmin() {
		return h.eventBus.Dispatch(&events.UserDeleteFailedEvent{Id: command.Id, Err: errors.New("admin role cannot be deleted")})
	}

	aggregate, err := h.repo.GetUserAggregate(&command.Id)

	// 如果用户不存在，直接返回
	if err != nil {
		h.eventBus.Dispatch(&events.UserDeleteFailedEvent{Id: command.Id, Err: err})
		return errors.New("用户不存在")
	}

	tx := h.repo.Begin()
	// 删除聚合
	if err := h.repo.DeleteUserAggregate(aggregate.UserId); err != nil {
		tx.Rollback()
		return err
	}

	// 发布事件
	if err := h.eventBus.Dispatch(&events.UserDeletedEvent{Id: command.Id}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}