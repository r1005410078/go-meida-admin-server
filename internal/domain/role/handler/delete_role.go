package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/role"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type DeleteRoleCommandHandler struct {
	repo role.IRoleAggregateRepository
	IEventBus shared.IEventBus
}

func NewDeleteRoleCommandHandler(repo role.IRoleAggregateRepository, IEventBus shared.IEventBus) *DeleteRoleCommandHandler {
	return &DeleteRoleCommandHandler{
		repo,
		IEventBus,
	}
}

func (h *DeleteRoleCommandHandler) Handle(command *command.DeleteRoleCommand) error {
	aggregate, err := h.repo.GetRoleAggregate(command.Id)

	// 如果不是是管理员，不允许删除
	if !h.repo.IsAdmin() {
		return h.IEventBus.Dispatch(events.RoleDeleteFailedEvent{Id: command.Id, Err: errors.New("admin role cannot be deleted")})
	}

	// 如果角色不存在，直接返回
	if err != nil {
		h.IEventBus.Dispatch(events.RoleDeleteFailedEvent{Id: command.Id, Err:  errors.New("role not found")})
		return err
	}

	if err := aggregate.Delete(command, h.IEventBus); err != nil {
		return err
	}
	
	return nil
}