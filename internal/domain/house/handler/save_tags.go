package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/house"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)


type SaveTagCommandHandler struct {
	repo house.IHouseAggregateRepository
	eventBus shared.IEventBus
}

func NewSaveTagCommandHandler(repo house.IHouseAggregateRepository, eventBus shared.IEventBus) *SaveTagCommandHandler {
	return &SaveTagCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *SaveTagCommandHandler) Handle(command *command.SaveHouseTagsCommand) error {
	aggregate, err := h.repo.GetAggregate(command.HousePropertyID)
	if aggregate == nil || err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return errors.New("房源不存在")
	}

	if err := h.eventBus.Dispatch(&events.SaveHouseTagsEvent{ SaveHouseTagsCommand: command}); err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	return nil
}