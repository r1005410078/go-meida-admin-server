package repository

import (
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
)

// IUserRepository 用户仓储接口
type IUserRepository interface {
	FindById(id string) (*model.User, error)
	Save(user *model.User) error
	Delete(user *model.User) error
	List() ([]model.User, error)
	// 关联角色
	AssoicatedRoles(event *events.AssoicatedRolesEvent) error
	// 删除用户
	DeleteUser(event *events.UserDeletedEvent) error
	// 保存用户
	SaveUser(event *events.SaveUserEvent) error
	// 保存用户状态
	SaveUserStatus(user *events.UserStatusEvent) error
}