package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserHandler struct {
	repo user.IUserAggregateRepository
	eventBus  shared.IEventBus
}

func NewLoginUserHandler(repo user.IUserAggregateRepository, eventBus shared.IEventBus) *LoginUserHandler {
	return &LoginUserHandler{
		repo: repo,
		eventBus:  eventBus,
	}
}

func (h *LoginUserHandler) Handle(ctx *gin.Context, cmd command.LoginInCommand) error {
	// 获取聚合信息
	aggregate, err := h.repo.GetUserAggregateByUsername(cmd.Username)
	if err != nil {
		h.eventBus.Dispatch(&events.LoginFailedEvent{
			Username:  cmd.Username,
			Error:     "user not found",
			Timestamp: time.Now(),
			IP:        h.getClientIP(ctx),
			Attempts:  aggregate.GetAttempts(),
		})
		return err
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(*aggregate.PasswordHash), []byte(cmd.Password)); err != nil {
		// 密码错误
		aggregate.LoginFailed()
		h.eventBus.Dispatch(&events.LoginFailedEvent{
			Username:  cmd.Username,
			Error:     "invalid password",
			Timestamp: time.Now(),
			IP:        h.getClientIP(ctx),
			Attempts:  aggregate.GetAttempts(),
		})

		return err
	}
 
	// 检查用户状态
	if err := aggregate.CheckStatusActive(); err != nil {
		return h.eventBus.Dispatch(&events.LoginFailedEvent	{
			Username:  cmd.Username,
			Error:     "user is not active",
			Timestamp: time.Now(),
			IP:        h.getClientIP(ctx),
			Attempts:  aggregate.GetAttempts(),
		})
	}

	// 重置登录尝试次数
	aggregate.LoggedIn()
	
	// 保存聚合
	tx := h.repo.Begin()
	if err := h.repo.SaveUserAggregate(aggregate); err != nil {
		tx.Rollback()
		return err
	}

	// 发送登录成功事件
	if err := h.eventBus.Dispatch(&events.LoggedInEvent{
		ID:        aggregate.UserId,
		Username:  aggregate.Username,
		LoginTime: time.Now(),
	}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (h *LoginUserHandler) getClientIP(ctx *gin.Context) string {
	// TODO: 从上下文中获取客户端IP
	return ctx.ClientIP()
}

