package forms

type IFormsAggregateRepository interface {
	Begin()
	Commit()
	Rollback()
	
	Save(aggregate *FormAggregate) error
	Updates(aggregate []*FormAggregate) error
	GetAggregate(formId, fieldId *string) (*FormAggregate, error)
	GetAggregates(formId  *string) ([]*FormAggregate, error)
	DeleteAggregate(id *string)error
	// 字段名称重复
	ExistFieldName (formId, fieldName *string) bool
	// 表单名称重复
	ExistFormName (formName *string) bool
	// 字段ids重复
	ExistFieldIds(formId string, ids []string) bool
}
