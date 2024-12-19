package permissions

import "github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"


type PermissionsRepositoryer interface {
	// 保存权限
	Save(permission *Permission) error
	// 删除权限
	Delete(permission *Permission) error
	// 权限列表
	List() ([]*model.UserPermission, error)
}