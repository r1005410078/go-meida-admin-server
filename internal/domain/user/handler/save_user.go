package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
)

type SaveUserCommandHandler struct {
	repo user.IUserAggregateRepository
	eventBus shared.IEventBus
}

func NewSaveUserCommandHandler(repo user.IUserAggregateRepository, eventBus shared.IEventBus) *SaveUserCommandHandler {
	return &SaveUserCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *SaveUserCommandHandler) Handle(command *command.SaveUserCommand) error {
	// 如果不是是管理员，不允许保存
	if !h.repo.IsAdmin() {
		return h.eventBus.Dispatch(events.NewSaveUserFailedEvent(command.ToEvent(), errors.New("admin role cannot be saved")))
	}
	
	var aggregate *user.UserAggregate


	if command.ID == nil {
		// 用户名和手机号已经存在
		if h.repo.ExistUser(command.Username) {
			return h.eventBus.Dispatch(events.NewSaveUserFailedEvent(command.ToEvent(), errors.New("user already exists")))
		}
		
		var err error
		// 创建聚合
		aggregate, err = user.New(user.UserAggregate{
			Username: command.Username,
			Phone: command.Phone,
			Email: command.Email,
		})

		if err != nil {
			h.eventBus.Dispatch(events.NewSaveUserFailedEvent(command.ToEvent(), err))
			return err
		}
	} else {
		var err error
		aggregate, err = h.repo.GetUserAggregate(command.ID)
		if err != nil {
			return err
		}

		// 更新聚合
		if err := aggregate.Update(command, h.eventBus); err != nil {
			return err
		}
	}

	// 开事务保存用户聚合
	h.repo.Begin()
	if err := h.repo.SaveUserAggregate(aggregate); err != nil {
		h.repo.Rollback()
		return err
	}

	// 发布保存用户事件
	command.ID = aggregate.UserId
	if err := h.eventBus.Dispatch(command.ToEvent()); err != nil {
		h.repo.Rollback()
		return err
	}

	// 如果有状态变化，发布更新状态事件
	if command.Status != nil {
		if err := h.eventBus.Dispatch(&events.UserStatusEvent {Id: *aggregate.UserId, Status: command.Status}); err != nil {
			h.repo.Rollback()
			return err
		} 
	}

	// 如果有角色变化，发布更新角色事件
	if command.RoleId != nil  {
		if err := h.eventBus.Dispatch(&events.AssoicatedRolesEvent{UserId: *aggregate.UserId, RoleId: *command.RoleId }); err != nil {
			h.repo.Rollback()
			return err
		}
	}

	return h.repo.Commit()
}