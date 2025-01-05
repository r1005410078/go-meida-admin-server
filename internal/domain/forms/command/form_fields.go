package command

import "github.com/r1005410078/meida-admin-server/internal/domain/forms/values"

type SaveFormsFieldsCommand struct {
	FormID          *string                 `json:"form_id" binding:"required"` 	 // 表单 ID
	FiledId         *string                `json:"field_id" ` 										// 字段名称
	Label           *string                `json:"label"` 									 			// 字段名称
	Type            *string                `json:"type"` 													// text, select, radio, map.
	Required        *bool                  `json:"required"` 							 				// 是否必填
	Placeholder     *string                `json:"placeholder"`					 					// 占位符
	ValidationRules *map[string]string     `json:"validation_rules"` 							// 校验规则 (JSON 解析)
	Options         *map[string]string     `json:"options"`      									// 选项
	DependsOn       *[]*values.Dependency  `json:"dependencies"` 									// 关联联动规则
}

type DeleteFormsFieldsCommand struct {
	FormID         *string    `json:"form_id" binding:"required"`
	FiledId        *string    `json:"field_id" binding:"required"` 		
}

