package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
)

type AssoicatedRolesCommandHandler struct {
	repo user.IUserAggregateRepository
	eventBus shared.IEventBus
}

func NewAssoicatedRolesCommandHandler(repo user.IUserAggregateRepository, eventBus shared.IEventBus) *AssoicatedRolesCommandHandler {
	return &AssoicatedRolesCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *AssoicatedRolesCommandHandler) Handle(command *command.AssociatedRolesCommand) error {

	// 检查权限
	if !h.repo.IsAdmin() {
		return h.eventBus.Dispatch(events.NewAssoicatedRolesFailedEvent(command.ToEvent(), errors.New("没有权限")))
	}

	// 检查角色是否存在
	if !h.repo.ExistRole(command.RoleId) {
		return h.eventBus.Dispatch(events.NewAssoicatedRolesFailedEvent(command.ToEvent(), errors.New("角色不存在")))
	}
	
	aggregate, err := h.repo.GetUserAggregate(command.UserId)
	// 聚合不存在，直接返回
	if err != nil {
		return err
	}

	// 角色已关联，直接返回
	if err := aggregate.AssociatedRoles(command, h.eventBus); err != nil {
		return err
	}

	tx := h.repo.Begin()
	// 开启事务保存聚合
	if err := h.repo.SaveUserAggregate(aggregate); err != nil {
		tx.Rollback()
		return err
	}

	// 发布事件
	if err := h.eventBus.Dispatch(command.ToEvent()); err != nil {
		return err
	}

	return tx.Commit().Error
}