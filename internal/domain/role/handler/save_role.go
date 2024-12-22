package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/role"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type SaveRoleCommandHandler struct {
	repo role.IRoleAggregateRepository
	IEventBus shared.IEventBus
}

func NewSaveRoleCommandHandler(repo role.IRoleAggregateRepository, eventBus shared.IEventBus) *SaveRoleCommandHandler {
	return &SaveRoleCommandHandler{
		repo: repo,
		IEventBus: eventBus,
	}
}

func (h *SaveRoleCommandHandler) Handle(command *command.SaveRoleCommand) error {
	// 不是管理员
	if !h.repo.IsAdmin() {
		return h.IEventBus.Dispatch(events.NewRoleSaveFailedEvent(*command.ToEvent(), errors.New("没有权限")))
	}

	// 角色名称已经存在
	if h.repo.IsRoleNameExist(command.Name) {
		return h.IEventBus.Dispatch(events.NewRoleSaveFailedEvent(*command.ToEvent(), errors.New("角色名称已存在")))
	}

	// 角色权限id无效
	if !h.repo.ExistsPermissionIds(command.PermissionIds) {
		return h.IEventBus.Dispatch(events.NewRoleSaveFailedEvent(*command.ToEvent(), errors.New("权限id已存在")))
	}

	var aggregate *role.RoleAggregate
	// 如果没有id，创建新的聚合
	if command.Id == nil {
	  aggregate =	role.NewRoleAggregate(command.Name)
	} else {
		// 如果有id，获取聚合
		var err error
		aggregate, err = h.repo.GetRoleAggregate(*command.Id)
		if err != nil {
			return h.IEventBus.Dispatch(events.NewRoleSaveFailedEvent(*command.ToEvent(), errors.New("没找到角色")))
		}
	}

	// 保存聚合
	if err := aggregate.Save(command, h.IEventBus); err != nil {
		return err
	}

	tx := h.repo.Begin()
	if err := h.repo.SaveRoleAggregate(aggregate); err != nil {
		tx.Rollback()
		return err
	}

	// 触发角色保存成功事件
	command.Id = aggregate.Id
	if err := h.IEventBus.Dispatch(*command.ToEvent()); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		
		return err
	}

	return nil
}

