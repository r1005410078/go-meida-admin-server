package handler

import (
	"errors"
	"slices"

	"github.com/r1005410078/meida-admin-server/internal/domain/forms"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/command"
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

func (h *SaveDependsOnCommandHandler) Handle(cmd *command.SaveDependsOnCommand) error {
	// 获取聚合
	aggregate, err := h.repo.GetAggregate(&cmd.FormID)
	if err != nil {
		return errors.New("关联失败，表单字段不存在")
	}

	fields := aggregate.Fields

	// 判断字段是否存在
	existField := slices.ContainsFunc(fields, func(v forms.FormField) bool { return *v.FieldId == cmd.FieldID })
	if !existField {
		return errors.New("关联失败，表单字段不存在")
	}

	// 判断联动字段是否存在
	for _, depend := range cmd.DependsOn {
		existDependField := slices.ContainsFunc(fields, func(v forms.FormField) bool { return *v.FieldId == depend.FieldId })
		if !existDependField {
			return errors.New("关联失败，联动字段不存在")
		}
	}

	h.repo.Begin()
	defer h.repo.Commit()

	event, err := aggregate.SaveRelatedIds(cmd)
	if err != nil {
		return err
	}

	// 保存聚合
	if err := h.repo.SaveAggregate(aggregate); err != nil {
		h.repo.Rollback()
		return err
	}

	// 发布事件
	if err := h.busEvent.Dispatch(event); err != nil {
		h.repo.Rollback()
		return err
	}

	return nil
}