package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/house"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type DeleteHouseCommandHandler struct {
	repo house.IHouseAggregateRepository
	eventBus shared.IEventBus
}

func NewDeleteHouseCommandHandler(repo house.IHouseAggregateRepository, eventBus shared.IEventBus) *DeleteHouseCommandHandler {
	return &DeleteHouseCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *DeleteHouseCommandHandler) Handle(command *command.DeleteHouseCommand) error {
	
	aggregate, err := h.repo.GetAggregate(command.ID)
	if err != nil {
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return errors.New("房源不存在")
	}

	deleteEvent := aggregate.Delete(command)

	h.repo.Begin()
	defer h.repo.Commit()
	if err := h.repo.DeleteAggregate(command.ID); err != nil {
		return err
	}

	if err := h.eventBus.Dispatch(deleteEvent); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(events.NewHouseCommandError(command, err))
		return err
	}

	return nil
}