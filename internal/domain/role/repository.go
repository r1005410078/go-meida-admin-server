package role

import "gorm.io/gorm"

type IRoleAggregateRepository interface {
	// 开启事务
	Begin() *gorm.DB

	// 是否是管理员
	IsAdmin() bool
	// 权限id是否有效
	IsValidPermissionID(id string) bool
	// 保存聚合
	SaveRoleAggregate(aggregate *RoleAggregate) error
	// 删除聚合
	DeleteRoleAggregate(id string) error
	// 获取聚合
	GetRoleAggregate(id string) (*RoleAggregate, error)
	// 是否有权限
	ExistsPermissionIds(ids []string) bool
	// 角色名称是否存在
	IsRoleNameExist(name string) bool
}

