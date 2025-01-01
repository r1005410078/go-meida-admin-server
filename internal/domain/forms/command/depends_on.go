package command

import "github.com/r1005410078/meida-admin-server/internal/domain/forms/values"

// 关联字段
type SaveDependsOnCommand struct {
	FormID         string     `json:"form_id" binding:"required"` // 表单 ID
	FieldID         string     `json:"field_id" binding:"required"`
	DependsOn 		 []*values.Dependency `json:"dependsOn" binding:"required"`
}
