package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/forms"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type SaveFormsHandler struct {
	repo forms.IFormsAggregateRepository 
	eventBus shared.IEventBus
}

func NewSaveFormsHandler(repo forms.IFormsAggregateRepository, eventBus shared.IEventBus) *SaveFormsHandler {
	return &SaveFormsHandler{
		repo: repo,
		eventBus: eventBus,
	}
}	

func (h *SaveFormsHandler) Handle(command *command.SaveFormsCommand) error {
	// 字段名称重复
	if h.repo.ExistFormName(&command.Name) {
		h.eventBus.Dispatch(&events.SaveFormsFailedEvent{SaveFormsCommand: command, Err: errors.New("字段名称重复")})
		return errors.New("字段名称重复")
	}
		
  var aggregate *forms.FormAggregate

	h.repo.Begin()
	defer h.repo.Commit()
	
	if command.Id == nil {
		aggregate = forms.New(aggregate)
		// 保存聚合
		if err := h.repo.Save(aggregate); err != nil {
			h.repo.Rollback()
			// 保存失败
			h.eventBus.Dispatch(events.SaveFormsFailedEvent { SaveFormsCommand: command, Err: err })
			return err
		}

		// 新增表单
		command.Id = &aggregate.FormId
		if err := h.eventBus.Dispatch(events.CreateFormsEvent { SaveFormsCommand: command }); err != nil {
			h.repo.Rollback()
			return err
		}

		return nil
	}

	// 更新名称
	aggregates, err := h.repo.GetAggregates(command.Id)
	if err != nil {
		// 保存失败
		h.eventBus.Dispatch(events.SaveFormsFailedEvent { SaveFormsCommand: command, Err: err })
		return err
	}

	if err := h.repo.Updates(aggregates); err != nil {
		h.repo.Rollback()
		// 保存失败
		h.eventBus.Dispatch(events.SaveFormsFailedEvent { SaveFormsCommand: command, Err: err })
		return err
	}

	// 更新表单
	if err := h.eventBus.Dispatch(events.UpdateFormsEvent { SaveFormsCommand: command }); err != nil {	
		h.repo.Rollback()
		return err
	}

	return nil
}