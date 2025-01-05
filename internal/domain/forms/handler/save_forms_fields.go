package handler

import (
	"github.com/r1005410078/meida-admin-server/internal/domain/forms"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type SaveFormsFieldsHandler struct {
	repo forms.IFormsAggregateRepository
	eventBus shared.IEventBus
}

func NewSaveFormsFieldsHandler(repo forms.IFormsAggregateRepository, eventBus shared.IEventBus) *SaveFormsFieldsHandler {
	return &SaveFormsFieldsHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *SaveFormsFieldsHandler) Handle(cmd *command.SaveFormsFieldsCommand) error {
	// 更新字段
	aggregate, err := h.repo.GetAggregate(cmd.FormID)
	if err != nil {
		h.eventBus.Dispatch(&events.UpdateFormsFieldsFailedEvent{SaveFormsFieldsCommand: cmd, Err: err})
		return err
	}

 	event, err := aggregate.SaveFormsFields(cmd) 
	if err != nil {
		h.eventBus.Dispatch(&events.UpdateFormsFieldsFailedEvent{SaveFormsFieldsCommand: cmd, Err: err})
		return err
	}
	
	// 更新聚合
	h.repo.Begin()
	defer h.repo.Commit()
	
	// 保存聚合
	if err := h.repo.SaveAggregate(aggregate); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(&events.UpdateFormsFieldsFailedEvent{SaveFormsFieldsCommand: cmd, Err: err})
		return err
	}	

	// 发布保存事件
	if err := h.eventBus.Dispatch(event); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(&events.UpdateFormsFieldsFailedEvent{SaveFormsFieldsCommand: cmd, Err: err})
		return err
	}

	return nil
}