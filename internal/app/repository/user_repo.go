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
	// 根据邮箱获取用户
	FindUserByEmail(email string) (*model.User, error)
	// 保存Email验证码
	SaveEmailCode(email string, code string) error
	// 获取验证码
	GetEmailCode(email *string) *string
	// 保存登陆token
	SaveLoginToken(userId *string) error
	// 删除登陆token
	DeleteLoginToken(userId *string) error
}