package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
)


type LogoutUserHandler struct {
	repo user.IUserAggregateRepository
	eventBus shared.IEventBus
}

func NewLogoutUserHandler(repo user.IUserAggregateRepository, eventBus shared.IEventBus) *LogoutUserHandler {
	return &LogoutUserHandler{
		repo: repo,
		eventBus:  eventBus,
	}
}

func (h *LogoutUserHandler) Handle(ctx *gin.Context, cmd command.LoggedOutCommand) error {

	aggregate, err := h.repo.GetUserAggregate(&cmd.UserId)
	if err != nil {
		return err
	}

	if aggregate == nil {
		err := errors.New("用户不存在")
		h.eventBus.Dispatch(&events.LoggedOutFailedEvent{Token: cmd.Token, UserId: cmd.UserId, Err: err})
		// 聚合不存在，退出失败
		return err
	}

	// 更新聚合
	aggregate.LoggedOut()

	// 保存聚合
	h.repo.Begin()
	if err := h.repo.SaveUserAggregate(aggregate); err != nil {
		h.repo.Rollback()
		return err
	}

	// 发布事件
	if err := h.eventBus.Dispatch(&events.LoggedOutEvent{UserId: *aggregate.UserId}); err != nil {
		h.repo.Rollback()
		return err
	}

	return h.repo.Commit()
}