package user

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
	"gorm.io/gorm"
)

type UserAggregate struct {
	UserId       *string
	Username         *string
	Phone        *string
	Status       *string
	Role         *string
	PasswordHash *string
	DeletedAt    *gorm.DeletedAt
}


func NewUserAggregate(username *string, phone *string, status *string) *UserAggregate {
	return &UserAggregate{
		UserId:    shared.NewId(),
		Username:  username,
		Phone:     phone,
		Status:    status,
		DeletedAt: nil,
	}
}

func (u *UserAggregate) Update(command *command.SaveUserCommand, bus shared.IEventBus) error {
	// 验证密码负责度
	// TODO
	u.Username = command.Username
	u.Phone = command.Phone
	u.Status = command.Status
	u.Role = command.RoleId
	u.PasswordHash = command.PasswordHash

	return nil
}

func (u *UserAggregate) SaveStatus(command *command.UserStatusCommand, eventBus shared.IEventBus) error {
	if *command.Status == *u.Status {
		return eventBus.Dispatch(events.NewUserStatusFailedEvent(command.ToEvent(), errors.New("不能重复设置同样的状态")))
	}

	u.Status = command.Status
	return nil
}

func (u *UserAggregate) AssociatedRoles(command *command.AssociatedRolesCommand, bus shared.IEventBus) error {
	if u.Role != nil && command.RoleId == *u.Role {
		return bus.Dispatch(events.NewAssoicatedRolesFailedEvent(command.ToEvent(), errors.New("角色重复被关联")))
	}

	u.Role = &command.RoleId
	return nil
}
