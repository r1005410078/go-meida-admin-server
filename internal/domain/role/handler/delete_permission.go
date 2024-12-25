package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/role"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)


type DeletePermissionHandler struct {
	IEventBus shared.IEventBus
	repo role.IRoleAggregateRepository
}

func NewDeletePermissionHandler(repo role.IRoleAggregateRepository, IEventBus shared.IEventBus) *DeletePermissionHandler {
	return &DeletePermissionHandler{
		IEventBus,
		repo,
	}
}

func (h *DeletePermissionHandler) Handle(command *command.DeletePermissionCommand) error {
	// 不是管理员，没有权限
	if !h.repo.IsAdmin() {
		return h.IEventBus.Dispatch(events.NewDeletePermissionFailedEvent(command.RoleId, command.PermissionIds, errors.New("no permission")))
	}

	aggregate, err := h.repo.GetRoleAggregate(command.RoleId)

	// 如果没有找到角色，直接返回
	if err != nil || aggregate == nil {
		return h.IEventBus.Dispatch(events.NewDeletePermissionFailedEvent(command.RoleId, command.PermissionIds, errors.New("role not found")))
	}

	// 是否存在权限id
	if !h.repo.ExistsPermissionIds(command.PermissionIds) {
		return h.IEventBus.Dispatch(events.NewDeletePermissionFailedEvent(command.RoleId, command.PermissionIds, errors.New("permission not found")))
	}

	if err := aggregate.DeletePermission(command, h.IEventBus); err != nil {
		return err
	}

	return nil
}


