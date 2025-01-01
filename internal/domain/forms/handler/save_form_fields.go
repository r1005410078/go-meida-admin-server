package handler

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/forms"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type SaveFormsFiledsHandler struct {
	repo forms.IFormsAggregateRepository
	eventBus shared.IEventBus
}

func NewSaveFormsFiledsHandler(repo forms.IFormsAggregateRepository, eventBus shared.IEventBus) *SaveFormsFiledsHandler {
	return &SaveFormsFiledsHandler{
		repo: repo,
		eventBus: eventBus,
	}
}

func (h *SaveFormsFiledsHandler) Handle(cmd *command.SaveFormsFiledsCommand) error {
	// 字段名称重复
	if h.repo.ExistFieldName(&cmd.FormID, cmd.Label) {
		h.eventBus.Dispatch(&events.UpdateFormsFiledsFailedEvent{SaveFormsFiledsCommand: cmd, Err: errors.New("字段名称重复")})
		return errors.New("字段名称重复")
	}

	var aggregate *forms.FormAggregate
	if cmd.FiledId == nil {
		// 新增字段
		aggregate = forms.New(&forms.FormAggregate{
			FormId: cmd.FormID,
			FieldName: cmd.Label,
		})

		// 保存聚合
		h.repo.Begin()
		defer h.repo.Commit()
	
		// 报错聚合
		if err := h.repo.Save(aggregate); err != nil {
			h.repo.Rollback()
			h.eventBus.Dispatch(&events.UpdateFormsFiledsFailedEvent{SaveFormsFiledsCommand: cmd, Err: err})
			return err
		}

		// 发布新增表单事件
		if err := h.eventBus.Dispatch(&events.CreateFormsFiledsEvent{SaveFormsFiledsCommand: cmd}); err != nil {
			h.repo.Rollback()
			h.eventBus.Dispatch(&events.UpdateFormsFiledsFailedEvent{SaveFormsFiledsCommand: cmd, Err: err})
			return err
		}

		// 保存关联字段
		if cmd.DependsOn != nil && len(*cmd.DependsOn) > 0 {
			if err := h.eventBus.Dispatch(&events.SaveDependsOnEvent{
				SaveDependsOnCommand: &command.SaveDependsOnCommand{
					FormID: cmd.FormID,
					FieldID: *cmd.FiledId,
					DependsOn: *cmd.DependsOn,
				},
			}); err != nil {
				h.repo.Rollback()
				return errors.New("保存关联字段失败")
			}
		}

		return nil
	}

	// 更新字段
	aggregate, err := h.repo.GetAggregate(&cmd.FormID, cmd.FiledId)

	if err != nil {
		h.eventBus.Dispatch(&events.UpdateFormsFiledsFailedEvent{SaveFormsFiledsCommand: cmd, Err: err})
		return errors.New("字段表单不存在")
	}
	
	if cmd.Label != nil {
		aggregate.SetFieldName(cmd.Label)
	}

	// 更新聚合
	h.repo.Begin()
	defer h.repo.Commit()
	
	// 报错聚合
	if err := h.repo.Save(aggregate); err != nil {
		h.repo.Rollback()
		h.eventBus.Dispatch(&events.UpdateFormsFiledsFailedEvent{SaveFormsFiledsCommand: cmd, Err: err})
		return err
	}	

	return nil
}