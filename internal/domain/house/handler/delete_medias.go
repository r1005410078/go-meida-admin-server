package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/house"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type DeleteHouseMediasCommandHandler struct {
	repo house.IHouseAggregateRepository
	eventBus shared.IEventBus
}

func NewDeleteHouseMediasCommandHandler(repo house.IHouseAggregateRepository, eventBus shared.IEventBus) *DeleteHouseMediasCommandHandler {
	return &DeleteHouseMediasCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}


func (h *DeleteHouseMediasCommandHandler) Handle(command *command.DeleteHouseMediasCommand) error {
	
	aggregate, err := h.repo.GetAggregate(command.ID)
	if err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return errors.New("房源不存在")
	}

	deleteEvent, err := aggregate.DeleteMedias(command)
	if err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return errors.New("删除房源媒体失败")
	}

	h.repo.Begin()
	defer h.repo.Commit()
	if err := h.repo.SaveAggregate(aggregate); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	if err := h.eventBus.Dispatch(deleteEvent); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	return nil
}