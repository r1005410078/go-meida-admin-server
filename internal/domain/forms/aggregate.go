package forms

import (
	"errors"
	"slices"
	"time"

	"github.com/r1005410078/meida-admin-server/internal/domain/forms/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

type FormAggregate struct {
	FormId    string
	FormName  string
	Fields    []FormField
	// 关联ids
	RelatedIds []string
}

type FormField struct {
	FieldId   *string
	FieldName *string
	CreateAt	*time.Time
	UpdateAt 	*time.Time
	DeleteAt  *time.Time
}

func New(aggregate *FormAggregate) *FormAggregate {
	newAgg := &FormAggregate{
		FormId:    aggregate.FormId,
		FormName:  aggregate.FormName,
	}

	if aggregate.FormId == "" {
		newAgg.FormId = *shared.NewId()
	}

	return newAgg
}

func (h *FormAggregate) SaveRelatedIds(command *command.SaveDependsOnCommand) (interface{}, error) {
	if command.Id == nil {
		command.Id = shared.NewId()
		h.RelatedIds = append(h.RelatedIds, *command.Id)

		return &events.CreateDependsOnEvent{SaveDependsOnCommand: command}, nil
	}

	if !slices.Contains(h.RelatedIds, *command.Id) {
		return nil, errors.New("id 不存在")
	}

	return &events.UpdateDependsOnEvent{SaveDependsOnCommand: command}, nil
}

func (h *FormAggregate) SaveFormsFields(cmd *command.SaveFormsFieldsCommand) (interface{}, error) {
	now := time.Now()
	if cmd.FiledId == nil {
		// 判断字段名称是否重复
		existName := slices.ContainsFunc(h.Fields, func(v FormField) bool { return *v.FieldName == *cmd.Label })
		if existName {
			return  nil, errors.New("字段名称重复")
		}
		// 新增字段
		formField := FormField{
			FieldId:   shared.NewId(),
			FieldName: cmd.Label,
			CreateAt:  &now,
		}

		cmd.FiledId = formField.FieldId
		h.Fields = append(h.Fields, formField)
		return &events.CreateFormsFieldsEvent{SaveFormsFieldsCommand: cmd}, nil
	}

	// 更新字段
	existNotId := slices.ContainsFunc(h.Fields, func(v FormField) bool { return *v.FieldId == *cmd.FiledId })
	if !existNotId {
		return nil, errors.New("FiledId 不存在")
	}

	for i, v := range h.Fields {
		if *v.FieldId == *cmd.FiledId {
			h.Fields[i].FieldName = cmd.Label
			h.Fields[i].UpdateAt = &now
			break
		}
	}

	return &events.UpdateFormsFieldsEvent{SaveFormsFieldsCommand: cmd}, nil
}

func (h *FormAggregate) DeleteFormsFields(cmd *command.DeleteFormsFieldsCommand) (interface{}, error) {
	// 更新字段
	existNotId := slices.ContainsFunc(h.Fields, func(v FormField) bool { return *v.FieldId == *cmd.FiledId })
	if !existNotId {
		return nil, errors.New("FiledId 不存在")
	}

	now := time.Now()
	for i, v := range h.Fields {
		if *v.FieldId == *cmd.FiledId {
			h.Fields[i].DeleteAt = &now
			break
		}
	}

	return &events.DeleteFormsFiledsEvent{DeleteFormsFieldsCommand: cmd}, nil
}


func (h *FormAggregate) UpdateByFormsCommand(command *command.SaveFormsCommand) {
	h.FormName = command.Name
}