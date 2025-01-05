package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/forms"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type DeleteFormsCommandHandler struct {
	repo forms.IFormsAggregateRepository
	eventBus shared.IEventBus
}

func NewDeleteFormsCommandHandler(repo forms.IFormsAggregateRepository, eventBus shared.IEventBus) *DeleteFormsCommandHandler {
	return &DeleteFormsCommandHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *DeleteFormsCommandHandler) Handle(command *command.DeleteFormsCommand) error {
	h.repo.Begin()
	defer h.repo.Commit()

	if err := h.repo.DeleteAggregateByFormId(&command.Id); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(&events.DeleteFormsFailedEvent { DeleteFormsCommand: command, Err: err })
		return errors.New("删除表单失败")
	}
	
	if err := h.eventBus.Dispatch(&events.DeleteFormsEvent { Id: &command.Id }); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(&events.DeleteFormsFailedEvent { DeleteFormsCommand: command, Err: err })
		return errors.New("删除表单失败")
	}

	return nil
}

