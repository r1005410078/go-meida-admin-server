package handler

import (
	"github.com/r1005410078/meida-admin-server/internal/domain/forms"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type DeleteFormsFiledsHandler struct {
	repo forms.IFormsAggregateRepository
	eventBus shared.IEventBus
}

func NewDeleteFormsFiledsHandler(repo forms.IFormsAggregateRepository, eventBus shared.IEventBus) *DeleteFormsFiledsHandler {
	return &DeleteFormsFiledsHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *DeleteFormsFiledsHandler) Handle(cmd *command.DeleteFormsFieldsCommand) error {
	h.repo.Begin()
	defer h.repo.Commit()

	aggregate, err := h.repo.GetAggregate(cmd.FormID)
	if err != nil {
		h.eventBus.Dispatch(&events.DeleteFormsFiledsFailedEvent{DeleteFormsFieldsCommand: cmd, Err: err})
		return err
	}

  event, err :=	aggregate.DeleteFormsFields(cmd)
	if err != nil {
		h.eventBus.Dispatch(&events.DeleteFormsFiledsFailedEvent{DeleteFormsFieldsCommand: cmd, Err: err})
		return err
	}

	// 保存聚合
	if err := h.repo.SaveAggregate(aggregate); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(&events.DeleteFormsFiledsFailedEvent{DeleteFormsFieldsCommand: cmd, Err: err})
		return err
	}
	
	// 发布事件
	if err := h.eventBus.Dispatch(event); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(&events.DeleteFormsFiledsFailedEvent{DeleteFormsFieldsCommand: cmd, Err: err})	
		return err
	}

	return nil
}