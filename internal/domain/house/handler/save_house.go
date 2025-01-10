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

	if command.ID == nil {
		// 新建房源
		aggreagte := house.New(command)
		
		h.repo.Begin()
		defer h.repo.Commit()
		if err := h.repo.SaveAggregate(aggreagte); err != nil {
			return err
		}

		// 发布新建房源事件
		if err := h.eventBus.Dispatch(events.CreateHouseEvent { SaveHouseCommand: command }); err != nil {
			h.repo.Rollback()
			h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
			return err
		}

		return nil
	}

	// 检查参数
	if h.repo.ExistAddress(command.ToAddress()) {
		return errors.New("房源地址已存在")
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
	updateEvent := aggregate.Update(command)
 
	// 保存聚合
	if err := h.repo.SaveAggregate(aggregate); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	// 发布更新房源事件
	if err := h.eventBus.Dispatch(updateEvent); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	return nil
}