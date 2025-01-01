package forms

import (
	"slices"
	"time"

	"github.com/r1005410078/meida-admin-server/internal/domain/forms/values"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)


type FormAggregate struct {
	FormId string 
	FormName string
	FieldId string
	FieldName *string
	DependsOn []*values.Dependency
	DeleteAt time.Time
}

func New(aggregate *FormAggregate) *FormAggregate {
	newAgg := &FormAggregate{
		FormId: aggregate.FormId,
		FieldId: aggregate.FieldId,
		DependsOn: []*values.Dependency{},
	}

	if aggregate.FormId == "" {
		newAgg.FormId = *shared.NewId()
	}

	if aggregate.FieldId == "" {
		newAgg.FieldId = *shared.NewId()
	}

	return newAgg
}

// 设置字段名称
func (h *FormAggregate) SetFieldName(name *string) {
	h.FieldName = name
}

// 设置字段依赖
func (h *FormAggregate) ResetDependsOn(dependsOn []*values.Dependency, busEvent *shared.IEventBus) ([]*values.Dependency, []*values.Dependency) {
	// 找出需要删除的依赖
	var deleteDependsOn []*values.Dependency
	for _, depend := range h.DependsOn {
		if !slices.Contains(dependsOn, depend) {
			deleteDependsOn = append(deleteDependsOn, depend)
		}
	}

	// 需要添加的依赖
	var appendDependsOn []*values.Dependency
	for _, depend := range dependsOn {
		if !slices.Contains(h.DependsOn, depend) {
			appendDependsOn = append(appendDependsOn, depend)
		}
	}

	h.DependsOn = dependsOn

	return deleteDependsOn, appendDependsOn
}