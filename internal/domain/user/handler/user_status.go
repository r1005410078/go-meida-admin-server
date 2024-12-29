package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
)

type UserStatusCommandHandler struct {
	repo user.IUserAggregateRepository
	eventBus shared.IEventBus
}

func NewUserStatusCommandHandler(repo user.IUserAggregateRepository, eventBus shared.IEventBus) *UserStatusCommandHandler {
	return &UserStatusCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *UserStatusCommandHandler) Handle(command *command.UserStatusCommand) error {

	// 如果不是管理员，不允许保存
	if !h.repo.IsAdmin() {
		h.eventBus.Dispatch(events.NewUserStatusFailedEvent(command.ToEvent(), errors.New("没有权限")))
		return errors.New("没有权限")
	}

	aggregate, err := h.repo.GetUserAggregate(&command.Id)
	if err != nil {
		// 用户不存在
		h.eventBus.Dispatch(events.NewUserStatusFailedEvent(command.ToEvent(), errors.New("用户不存在")))
		return err
	}

	if err := aggregate.SaveStatus(command, h.eventBus); err != nil {
		return err
	}

	// 如果有状态变化，保存聚合
	h.repo.Begin()
	if err := h.repo.SaveUserAggregate(aggregate); err != nil {
		h.repo.Rollback()
		return err
	}

	// 发布事件
	if err := h.eventBus.Dispatch(command.ToEvent()); err != nil {
		h.repo.Rollback()
		return err
	}

	return h.repo.Commit()
}