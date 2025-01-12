package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/house"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type SaveHouseMediasCommandHandler struct {
	repo house.IHouseAggregateRepository
	eventBus shared.IEventBus
}

func NewSaveHouseMediasCommandHandler(repo house.IHouseAggregateRepository, eventBus shared.IEventBus) *SaveHouseMediasCommandHandler {
	return &SaveHouseMediasCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *SaveHouseMediasCommandHandler) Handle(command *command.SaveHouseMediasCommand) error {

	aggregate, err := h.repo.GetAggregate(command.HousePropertyID)
	
	if aggregate == nil || err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return errors.New("房源不存在")
	}

	if err := h.eventBus.Dispatch(&events.SaveHouseMediasEvent{ SaveHouseMediasCommand: command}); err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}
	
	return nil
}