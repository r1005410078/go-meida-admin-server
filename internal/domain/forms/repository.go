package forms

import "github.com/r1005410078/meida-admin-server/internal/domain/forms/command"

type IFormsAggregateRepository interface {
	Begin()
	Commit()
	Rollback()
	
	SaveAggregate(aggregate *FormAggregate) error
	DeleteAggregateFields(cmd *command.DeleteFormsFieldsCommand) error
	Updates(aggregate *FormAggregate) error
	GetAggregate(formId  *string) (*FormAggregate, error)
	DeleteAggregateByFormId(formId *string)error
	DeleteAggregateByFieldId(formId *string)error
	// 字段名称重复
	ExistFieldName (formId, fieldName *string) bool
	// 表单名称重复
	ExistFormName (formName *string) bool
	// 字段ids重复
	ExistFieldIds(formId string, ids []string) bool
}
