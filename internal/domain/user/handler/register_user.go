package handler

import (
	"context"
	"errors"
	"time"

	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
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

func (h *RegisterUserHandler) Handle(ctx context.Context, cmd *command.RegisterCommand) error {

	// 创建用户
	aggregate, err := user.New(user.UserAggregate{
		Username:  &cmd.Username,
		Email:     &cmd.Email,
	})

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

	// 保存聚合
	h.repo.Begin()
	if err := h.repo.SaveUserAggregate(aggregate); err != nil {
		h.repo.Rollback()
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
		ID:        *aggregate.UserId,
		Username:  cmd.Username,
		Email:     cmd.Email,
		CreatedAt: time.Now(),
	}
	
	if err := h.bus.Dispatch(registeredEvent); err != nil {
		h.repo.Rollback()
		return err
	}

	return h.repo.Commit()
}

