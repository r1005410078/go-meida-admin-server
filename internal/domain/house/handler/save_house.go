package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/house"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)


type SaveHouseCommandHandler struct {
	repo house.IHouseAggregateRepository
	eventBus shared.IEventBus
}

func NewSaveHouseCommandHandler(repo house.IHouseAggregateRepository, eventBus shared.IEventBus) *SaveHouseCommandHandler {
	return &SaveHouseCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *SaveHouseCommandHandler) Handle(command *command.SaveHouseCommand) error {
	// 检查参数
	if h.repo.ExistAddress(command.ToAddress(), command.ID) {
		return errors.New("房源地址已存在")
	}

	if command.ID == nil {
		// 新建房源
		aggreagte, err := house.New(command)
		if err != nil {
			return err
		}
		
		h.repo.Begin()
		defer h.repo.Commit()
		if err := h.repo.SaveAggregate(aggreagte); err != nil {
			h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
			return err
		}

		command.ID = aggreagte.HousePropertyId
		// 发布新建房源事件
		if err := h.eventBus.Dispatch(&events.CreateHouseEvent { SaveHouseCommand: command }); err != nil {
			h.repo.Rollback()
			h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
			return err
		}

		// 保存tags
		if err := h.saveTags(command); err != nil {
			return err
		}

		// 保存medias
		if err := h.saveMedias(command); err != nil {
			return err
		}

		// 保存location
		if err := h.saveLocation(command); err != nil {
			return err
		}

		return nil
	}

	// 更新房源
	aggregate, err := h.repo.GetAggregate(command.ID)
	if err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	h.repo.Begin()
	defer h.repo.Commit()

	// 更新聚合
	updateEvent, err := aggregate.Update(command)
	if err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}
 
	// 保存聚合
	if err := h.repo.SaveAggregate(aggregate); err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	// 发布更新房源事件
	if err := h.eventBus.Dispatch(updateEvent); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	// 保存tags
	if err := h.saveTags(command); err != nil {
		return err
	}

	// 保存medias
	if err := h.saveMedias(command); err != nil {
		return err
	}

	// 保存location
	if err := h.saveLocation(command); err != nil {
		return err
	}

	return nil
}

func (h *SaveHouseCommandHandler) saveTags(command *command.SaveHouseCommand) error {
	// 保存tags
	if command.HouseDetails.Tags != nil {
		if err := h.eventBus.Dispatch(&events.SaveHouseTagsEvent{ SaveHouseTagsCommand: command.ToSaveHouseTagsCommand() }); err != nil {
			h.repo.Rollback()
			h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
			return err
		}
	}
	return nil
}

func (h *SaveHouseCommandHandler) saveMedias(command *command.SaveHouseCommand) error {
	// 保存medias
	if command.HouseDetails.Medias != nil {
		if err := h.eventBus.Dispatch(&events.SaveHouseMediasEvent{ SaveHouseMediasCommand: command.ToSaveHouseMediasCommand() }); err != nil {
			h.repo.Rollback()
			h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
			return err
		}
	}
	return nil
}

func (h *SaveHouseCommandHandler) saveLocation(command *command.SaveHouseCommand) error {
	// 保存location
	if command.HouseDetails.Location != nil {
		if err := h.eventBus.Dispatch(&events.SaveHouseLocationEvent{ SaveHouseLocationCommand: command.ToSaveHouseLocationCommand() }); err != nil {
			h.repo.Rollback()
			h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
			return err
		}
	}
	return nil
}