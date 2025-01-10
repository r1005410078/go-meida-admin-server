package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/house"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type SaveHouseLocationCommandHandler struct {
	repo house.IHouseAggregateRepository
	eventBus shared.IEventBus
}

func NewSaveHouseLocationCommandHandler(repo house.IHouseAggregateRepository, eventBus shared.IEventBus) *SaveHouseLocationCommandHandler {
	return &SaveHouseLocationCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *SaveHouseLocationCommandHandler) Handle(command *command.SaveHouseLocationCommand) error {

	aggregate, err := h.repo.GetAggregate(command.ID)
	if err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	if aggregate == nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, errors.New("房源不存在")))
		return errors.New("房源不存在")
	}

	if err := h.eventBus.Dispatch(events.SaveHouseLocationEvent{ SaveHouseLocationCommand: command }); err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	return nil
}