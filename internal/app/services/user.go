package services

import (
	"github.com/r1005410078/meida-admin-server/internal/app/repository"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"go.uber.org/zap"
)

type UserServices struct {
	repo repository.IUserRepository
	logger *zap.Logger
}

func NewUserServices(repo repository.IUserRepository, logger *zap.Logger) *UserServices {
	return &UserServices{
		repo,
		logger,
	}
}

func (u *UserServices) FindById(userId string) (*model.User, error) {
	return u.repo.FindById(userId)
}

func (u *UserServices) List() ([]model.User, error) {
	return u.repo.List()
}

// 关联角色
func (u *UserServices) AssoicatedRolesEventHandle(event *events.AssoicatedRolesEvent) error {
	return u.repo.AssoicatedRoles(event)
}

// 关联角色失败
func (u *UserServices) AssoicatedRolesFailedEventHandle(event *events.AssoicatedRolesFailedEvent) error {
	u.logger.Error(event.Err.Error())
	return event.Err
}

// 删除用户
func (u *UserServices) DeleteUserHandle(event *events.UserDeletedEvent) error {
	return u.repo.DeleteUser(event)
}

// 删除用户失败	
func (u *UserServices) DeleteUserFailedEventHandle(event *events.UserDeleteFailedEvent) error {
	u.logger.Error(event.Err.Error())
	return event.Err
}

// 保存用户
func (u *UserServices) SaveUserEventHandle(event *events.SaveUserEvent) error {
	return u.repo.SaveUser(event)
}

// 保存用户失败
func (u *UserServices) SaveUserFailedEventHandle(event *events.SaveUserFailedEvent) error {
	u.logger.Error(event.Err.Error())
	return event.Err
}

// 更改用户状态
func (u *UserServices) SaveUserStatusEventHandle(event *events.UserStatusEvent) error {
	return u.repo.SaveUserStatus(event)
}

// 更改用户状态失败
func (u *UserServices) SaveUserStatusFailedEventHandle(event *events.UserStatusFailedEvent) error {
	u.logger.Error(event.Err.Error())
	return event.Err
}