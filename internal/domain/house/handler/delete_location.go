package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/house"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type DeleteHouseLocationCommandHandler struct {
	repo house.IHouseAggregateRepository
	eventBus shared.IEventBus
}

func NewDeleteHouseLocationCommandHandler(repo house.IHouseAggregateRepository, eventBus shared.IEventBus) *DeleteHouseLocationCommandHandler {
	return &DeleteHouseLocationCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *DeleteHouseLocationCommandHandler) Handle(command *command.DeleteHouseLocationCommand) error {

	aggregate, err := h.repo.GetAggregate(command.ID)
	if err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	if aggregate == nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, errors.New("房源不存在")))
		return errors.New("房源不存在")
	}

	if err := h.eventBus.Dispatch(&events.DeleteHouseLocationEvent{ DeleteHouseLocationCommand: command }); err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	return nil
}