package user

import (
	"errors"
	"time"

	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
)

type UserAggregate struct {
	UserId    string
	Name      string
	Phone     string
	Status    string
	RoleId    string
	DeletedAt *time.Time
}



func NewUserAggregate(name string, phone string, status string) *UserAggregate {
	return &UserAggregate{
		UserId:    *shared.NewId(),
		Name:      name,
		Phone:     phone,
		Status:    status,
		DeletedAt: nil,
	}
}

func (u *UserAggregate) Update(command *command.SaveUserCommand, bus shared.IEventBus) error {
	if u.DeletedAt != nil {
		bus.Dispatch(events.NewSaveUserFailedEvent(command.ToEvent(), errors.New("用户已删除")))
		return errors.New("用户已删除")
	}

	// 验证密码负责度
	// TODO

	u.Name = command.Username
	u.Phone = command.Phone

	return nil
}

func (u *UserAggregate) SaveStatus(command *command.UserStatusCommand, eventBus shared.IEventBus) error {
	if u.DeletedAt != nil {
		eventBus.Dispatch(events.NewUserStatusFailedEvent(command.ToEvent(), errors.New("用户已删除")))
		return errors.New("用户已删除")
	}

	if command.Status == u.Status {
		eventBus.Dispatch(events.NewUserStatusFailedEvent(command.ToEvent(), errors.New("状态未变化")))
		return errors.New("状态更新失败")
	}

	u.Status = command.Status
	return nil
}


func (u *UserAggregate) AssociatedRoles(command *command.AssociatedRolesCommand, bus shared.IEventBus) error {
	if u.DeletedAt != nil {
		bus.Dispatch(events.NewAssoicatedRolesFailedEvent(command.ToEvent(), errors.New("用户已删除")))
		return errors.New("用户已删除")
	}

	if command.RoleId == u.RoleId {
		bus.Dispatch(events.NewAssoicatedRolesFailedEvent(command.ToEvent(), errors.New("角色已关联")))
		return errors.New("角色已关联")
	}

	u.RoleId = command.RoleId
	return nil
}