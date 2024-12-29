package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
)

type ChangePasswordHandler struct {
	repo    user.IUserAggregateRepository
	eventBus shared.IEventBus
}

func NewChangePasswordHandler(repo user.IUserAggregateRepository, eventBus shared.IEventBus) *ChangePasswordHandler {
	return &ChangePasswordHandler{
		repo:    repo,
		eventBus: eventBus,
	}
}

func (h *ChangePasswordHandler) Handle(command *command.ChangePasswordCommand) error {

	changePasswordEvent := events.ChangePasswordEvent{
		UserId:       command.UserId,
		PasswordHash: command.Password,
	}

	aggregate, err := h.repo.GetUserAggregate(&command.UserId)
	if err != nil {
		h.eventBus.Dispatch(events.NewChangePasswordFailedEvent(changePasswordEvent, err))
		return err
	}

	if aggregate == nil {
		err := errors.New("用户不存在")
		h.eventBus.Dispatch(events.NewChangePasswordFailedEvent(changePasswordEvent, err))
		// 聚合不存在，退出失败
		return err
	}

	
	if err := aggregate.ChangePassword(command); err != nil {
		h.eventBus.Dispatch(events.NewChangePasswordFailedEvent(changePasswordEvent, err))
		return err
	}

	// 保存聚合
	h.repo.Begin()
	if err := h.repo.SaveUserAggregate(aggregate); err != nil {
		h.repo.Rollback()
		return err
	}

	// 发布事件
	if err := h.eventBus.Dispatch(&changePasswordEvent); err != nil {
		h.repo.Rollback()
		return err
	}

	return h.repo.Commit()
}


