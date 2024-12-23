package repository

import (
	"github.com/r1005410078/meida-admin-server/internal/domain/role/events"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
)

type Roles struct {
	model.Role 
	Permissions []model.UserPermission `json:"permissions"`
}

type GetRoleListQuery struct {
	PageIndex int `json:"pageIndex" default:"1"`
	PageSize int`json:"pageSize" default:"20"`
}

type IRoleRepository interface {
	// 保存角色
	SaveRole(role  events.RoleSavedEvent) error
	// 删除角色
	DeleteRole(id string) error
	// 获取角色列表
	GetRoleList() ([]Roles, error) 
}