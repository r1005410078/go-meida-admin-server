package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/forms"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type SaveDependsOnCommandHandler struct {
	repo forms.IFormsAggregateRepository
	busEvent shared.IEventBus		
}

func NewSaveDependsOnCommandHandler(repo forms.IFormsAggregateRepository, busEvent shared.IEventBus) *SaveDependsOnCommandHandler {
	return &SaveDependsOnCommandHandler{
		repo: repo,
		busEvent: busEvent,
	}
}

func (h *SaveDependsOnCommandHandler) Handle(command *command.SaveDependsOnCommand) error {
	
  var dependsOnIds []string
	
	for _, v := range command.DependsOn {
		dependsOnIds = append(dependsOnIds, v.FieldId)
	}
	
	// 检查依赖表单是否存在
	if !h.repo.ExistFieldIds(command.FormID, dependsOnIds) {
		return errors.New("依赖表单不存在")
	}
	
	// 获取聚合
	aggregate, err := h.repo.GetAggregate(&command.FormID, &command.FieldID)
	if err != nil {
		return errors.New("关联失败，表单字段不存在")
	}
	
	deleteDependsOn, appendDependsOn := aggregate.ResetDependsOn(command.DependsOn, &h.busEvent)

	h.repo.Begin()
	defer h.repo.Commit()
	
	// 保存聚合
	if err := h.repo.Save(aggregate); err != nil {
		h.repo.Rollback()
		return err
	}
	
	if len(appendDependsOn) > 0 {
		if err := h.busEvent.Dispatch(&events.ApeendDependsOnEvent{SaveDependsOnCommand: command}); err != nil {
			h.repo.Rollback()
			return err
		}
	}
	
	if len(deleteDependsOn) > 0 {
		if err := h.busEvent.Dispatch(&events.DeleteDependsOnEvent{SaveDependsOnCommand: command}); err != nil {
			h.repo.Rollback()
			return err
		}
	}

	return nil
}