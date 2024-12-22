package command

import "github.com/r1005410078/meida-admin-server/internal/domain/role/events"

type SaveRoleCommand struct {
	// 角色ID
	Id *string `json:"id"`
	// 角色名称
	Name string `json:"name"`
	// 角色权限
	PermissionIds []string `json:"permissionIds"`
	// 角色描述
	Description string `json:"description"`
}

// 转化成事件
func (c *SaveRoleCommand) ToEvent() *events.RoleSavedEvent {
	return &events.RoleSavedEvent{
		Id:           c.Id,
		Name:         c.Name,
		PermissionIds: c.PermissionIds,
		Description:  c.Description,
	}
}