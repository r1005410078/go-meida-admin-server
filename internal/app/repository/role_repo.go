package repository

import "github.com/r1005410078/meida-admin-server/internal/domain/role/events"


type IRoleRepository interface {
	// 保存角色
	SaveRole(role  events.RoleSavedEvent) error
	// 删除角色
	DeleteRole(id string) error
}