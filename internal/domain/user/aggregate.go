package user

import (
	"errors"
	"strings"
	"time"

	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
	"golang.org/x/crypto/bcrypt"
)

type UserAggregate struct {
	UserId *string
	// 用户名
	Username *string
	// 手机
	Phone *string
	// 状态
	Status *string
	// 角色
	Role *string
	// 密码
	PasswordHash *string
	// 删除时间
	DeletedAt *time.Time
	// emil
	Email *string
	// 上次登录时间
	LastLoginAt *time.Time
	// 上次退出时间
	LastLogoutAt *time.Time
	// 登陆失败时间
	LoginFailedAt *time.Time
	// 登录尝试次数
	LoginAttempts *int32
}

func (h *UserAggregate) ChangePassword(command *command.ChangePasswordCommand) error {
	if err := checkPasswordStrength(command.Password); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(command.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	password := string(hashedPassword)
	h.PasswordHash = &password
	return nil
}

func New(agg UserAggregate) (*UserAggregate, error) {
	newAgg := &UserAggregate{
		UserId:        NewUserId(),
		Username:      agg.Username,
		Phone:         agg.Phone,
		Status:        agg.Status,
		Role:          agg.Role,
		PasswordHash:  agg.PasswordHash,
		Email:         agg.Email,
		LastLoginAt:   agg.LastLoginAt,
		LastLogoutAt:  agg.LastLogoutAt,
		LoginFailedAt: agg.LoginFailedAt,
		LoginAttempts: agg.LoginAttempts,
		DeletedAt:     nil,
	}

	if agg.PasswordHash != nil {
		if err := agg.ChangePassword(&command.ChangePasswordCommand{
			UserId:       *agg.UserId,
			Password:     *agg.PasswordHash,
		}); err != nil {
			return nil, err
		}
	}

	if agg.Status == nil {
		status := "active"
		agg.Status = &status
	}

	return newAgg, nil
}

func NewUserId() *string {
	// 更新删除事件
	return shared.NewId()
}

func (h *UserAggregate) LoggedOut() {
	// 更新退出事件
	now := time.Now()
	h.LastLogoutAt = &now
}

func (h *UserAggregate) LoggedIn() {
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

func (h *UserAggregate) GetLoginAttempts() int32 {
	if h.LoginAttempts == nil {
		return 0
	}
	return *h.LoginAttempts
}

// 实现增加登录尝试次数的逻辑
func (h *UserAggregate) IncrementLoginAttempts() int32 {
	if h.LoginAttempts == nil {
		h.LoginAttempts = new(int32)
	}
	*h.LoginAttempts += 1

	now := time.Now()
	h.LoginFailedAt = &now
	return h.GetLoginAttempts()
}

// 实现重置登录尝试次数的逻辑
func (h *UserAggregate) ResetLoginAttempts() {
	h.LoginAttempts = nil
}

// 检查用户状态
func (u *UserAggregate) CheckStatusActive() error {
	// 时间到了，重置尝试次数
	if u.GetLoginAttempts() > 0 && time.Now().After(u.LoginFailedAt.Add(24*time.Hour)) {
		u.ResetLoginAttempts()
	}

	if u.GetLoginAttempts() > 5 {
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

func (u *UserAggregate) Update(inputCommand *command.SaveUserCommand, bus shared.IEventBus) error {
	// 验证密码负责度
	// TODO
	u.Username = inputCommand.Username
	u.Phone = inputCommand.Phone
	u.Status = inputCommand.Status
	u.Role = inputCommand.RoleId

	if inputCommand.Password != nil {
		if err := u.ChangePassword(&command.ChangePasswordCommand{UserId: *u.UserId, Password: *inputCommand.Password}); err != nil {
			return err
		}
	}

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


/** 密码强弱检查 **/
func  checkPasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("密码长度必须至少为 8 个字符")
	}
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return errors.New("密码必须至少包含一个小写字母")
	}
	return nil
}
