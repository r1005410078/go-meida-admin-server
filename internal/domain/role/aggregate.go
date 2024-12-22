package role

import (
	"errors"
	"slices"
	"time"

	"github.com/r1005410078/meida-admin-server/internal/domain/role/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type RoleAggregate struct {
	Id        *string
	Name      string
	PermissionIds []string
	DeletedAt *time.Time
	UpdatedAt *time.Time
}

// 创建聚合
func NewRoleAggregate(name string) *RoleAggregate {
	return &RoleAggregate{
		Id: shared.NewId(),
		Name: name,
	}
}

// 更新聚合
func (r *RoleAggregate) Save(command *command.SaveRoleCommand, bus shared.IEventBus) error {
	// 如果已经删除了，直接返回
	if r.DeletedAt != nil {
		return bus.Dispatch(events.NewRoleSaveFailedEvent(*command.ToEvent(), errors.New("角色已经被删除")))
	}

	// 如果权限id已经存在，直接返回
	if r.isExistPermissionIds(command.PermissionIds) {
		return bus.Dispatch(events.NewRoleSaveFailedEvent(*command.ToEvent(), errors.New("权限id已存在")))
	}

	r.Name = command.Name
	now := time.Now()
	r.UpdatedAt = &now
	r.PermissionIds = append(r.PermissionIds, command.PermissionIds...)

	return nil
}

// 权限id已经存在
func (r *RoleAggregate) isExistPermissionIds(permissionIds []string) bool {
	for _, permissionId := range permissionIds {
		if slices.Contains(r.PermissionIds, permissionId) {
			return true
		}
	}
	return false
}

// 删除权限
func (r *RoleAggregate) DeletePermission(command *command.DeletePermissionCommand, bus shared.IEventBus) error {

	// 如果已经删除了，直接返回
	if r.DeletedAt != nil {
		return bus.Dispatch(events.NewDeletePermissionFailedEvent(command.Id, command.PermissionIds, errors.New("角色已经被删除")))
	}

	now := time.Now()

	r.UpdatedAt = &now
	
	// 删除权限
	for _, deletePermissionId := range command.PermissionIds {
		for i, permissionId := range r.PermissionIds {
			if deletePermissionId == permissionId {
				r.PermissionIds = append(r.PermissionIds[:i], r.PermissionIds[i+1:]...)
				break
			}
		}
	}
	
	// 触发角色删除成功事件
	return bus.Dispatch(events.DeletedPermissionEvent{Id: command.Id, PermissionId: command.PermissionIds})
}

// 删除角色
func (r *RoleAggregate) Delete(command *command.DeleteRoleCommand, IEventBus shared.IEventBus) error {
	// 如果已经删除了，直接返回
	if r.DeletedAt != nil {
		return IEventBus.Dispatch(events.RoleDeleteFailedEvent{Id: command.Id, Err: errors.New("role has been deleted")})
	}

	now := time.Now()
	r.DeletedAt = &now
	// 触发角色删除成功事件
	return IEventBus.Dispatch(events.RoleDeletedEvent{Id: command.Id})
}
