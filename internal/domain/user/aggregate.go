package user

import (
	"errors"
	"time"

	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
)

type UserAggregate struct {
	UserId       *string
	// 用户名
	Username     *string
	// 手机
	Phone        *string
	// 状态
	Status       *string
	// 角色
	Role         *string
	// 密码
	PasswordHash *string
	// 删除时间
	DeletedAt    *time.Time
	// emil
	Email        *string
	// 上次登录时间
	LastLoginAt  *time.Time
	// 上次退出时间
	LastLogoutAt *time.Time
	// 登陆失败时间
	LoginFailedAt *time.Time
	// 登录尝试次数
	Attempts     *int32	
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

func (h *UserAggregate) LoggedOut() {
	// 更新退出事件
	now := time.Now()
	h.LastLogoutAt = &now
}

func (h *UserAggregate) LoggedIn()  {
	// 更新登录事件
	now := time.Now()
	h.LastLoginAt = &now
	// 重置尝试次数
	h.ResetLoginAttempts()
}

// 登陆错误
func (h *UserAggregate) LoginFailed() {
	// 增加尝试次数
	h.IncrementLoginAttempts()
}

func (h *UserAggregate) GetAttempts() int32 {
	if h.Attempts == nil {
		return 0
	}
	return *h.Attempts
}

// 实现增加登录尝试次数的逻辑
func (h *UserAggregate) IncrementLoginAttempts() int32 {
	*h.Attempts += 1
	return h.GetAttempts()
}

// 实现重置登录尝试次数的逻辑
func (h *UserAggregate) ResetLoginAttempts() {
	h.Attempts = nil
}

// 检查用户状态
func (u *UserAggregate) CheckStatusActive() error {
	// 时间到了，重置尝试次数
	if u.GetAttempts() > 0 && time.Now().After(u.LoginFailedAt.Add(1*time.Hour)) {
		u.ResetLoginAttempts()
	}

	if u.GetAttempts() > 5 {
		return errors.New("login attempts exceeded")
	}

	if u.Status == nil {
		return nil
	}

	if *u.Status != "active" {
		return errors.New("user is not active")
	}

	return nil
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
