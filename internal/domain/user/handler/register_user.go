package handler

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserHandler struct {
	repo user.IUserAggregateRepository
	bus  shared.IEventBus
}

func NewRegisterUserHandler(repo user.IUserAggregateRepository, bus shared.IEventBus) *RegisterUserHandler {
	return &RegisterUserHandler{
		repo: repo,
		bus:  bus,
	}
}

func (h *RegisterUserHandler) Handle(ctx context.Context, cmd command.RegisterCommand) error {
	// 密码缺少强弱检查， 必须要有字母长度8位以上
	if err := h.checkPasswordStrength(cmd.Password); err != nil {
		failedEvent := &events.RegisterFailedEvent{
			Username:  cmd.Username,
			Email:     cmd.Email,
			Error:     err.Error(),
			Timestamp: time.Now(),
		}
		if err := h.bus.Dispatch(failedEvent); err != nil {
			return err
		}
		return err
	}

	// 验证邮箱验证码
	if err := h.repo.VerifyEmailCode(cmd.Email, cmd.VerificationCode); err != nil {
		failedEvent := &events.RegisterFailedEvent{
			Username:  cmd.Username,
			Email:     cmd.Email,
			Error:     err.Error(),
			Timestamp: time.Now(),
		}
		if err := h.bus.Dispatch(failedEvent); err != nil {
			return err
		}
		return err
	}

	// 检查用户名是否已存在
	if exists, err := h.repo.ExistsByUsername(cmd.Username); err != nil {
		failedEvent := &events.RegisterFailedEvent{
			Username:  cmd.Username,
			Email:     cmd.Email,
			Error:     err.Error(),
			Timestamp: time.Now(),
		}
		if err := h.bus.Dispatch(failedEvent); err != nil {
			return err
		}
		return err
	} else if exists {
		failedEvent := &events.RegisterFailedEvent{
			Username:  cmd.Username,
			Email:     cmd.Email,
			Error:     "username already exists",
			Timestamp: time.Now(),
		}
		if err := h.bus.Dispatch(failedEvent); err != nil {
			return err
		}
		return errors.New("username already exists")
	}

	// 检查邮箱是否已存在
	if exists, err := h.repo.ExistsByEmail(cmd.Email); err != nil {
		failedEvent := &events.RegisterFailedEvent{
			Username:  cmd.Username,
			Email:     cmd.Email,
			Error:     err.Error(),
			Timestamp: time.Now(),
		}
		if err := h.bus.Dispatch(failedEvent); err != nil {
			return err
		}
		return err
	} else if exists {
		failedEvent := &events.RegisterFailedEvent{
			Username:  cmd.Username,
			Email:     cmd.Email,
			Error:     "email already exists",
			Timestamp: time.Now(),
		}
		if err := h.bus.Dispatch(failedEvent); err != nil {
			return err
		}
		return errors.New("email already exists")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		failedEvent := &events.RegisterFailedEvent{
			Username:  cmd.Username,
			Email:     cmd.Email,
			Error:     err.Error(),
			Timestamp: time.Now(),
		}
		if err := h.bus.Dispatch(failedEvent); err != nil {
			return err
		}
		return err
	}

	// 创建用户
	hashedPasswordString := string(hashedPassword)
	aggregate := user.NewUserAggregate(&cmd.Username, &cmd.Email, &hashedPasswordString)
	if err := h.repo.SaveUserAggregate(aggregate); err != nil {
		failedEvent := &events.RegisterFailedEvent{
			Username:  cmd.Username,
			Email:     cmd.Email,
			Error:     err.Error(),
			Timestamp: time.Now(),
		}
		if err := h.bus.Dispatch(failedEvent); err != nil {
			return err
		}
		return err
	}

	// 发布注册成功事件
	registeredEvent := &events.RegisteredEvent{
		ID:        uuid.New().String(),
		Username:  cmd.Username,
		Email:     cmd.Email,
		CreatedAt: time.Now(),
	}
	
	if err := h.bus.Dispatch(registeredEvent); err != nil {
		return err
	}

	return nil
}

/** 密码强弱检查 **/
func (h *RegisterUserHandler) checkPasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("密码长度必须至少为 8 个字符")
	}
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return errors.New("密码必须至少包含一个小写字母")
	}
	return nil
}
