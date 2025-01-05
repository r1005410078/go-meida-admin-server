package values

type FieldOption struct {
	ID        int       `json:"id"` // 选项 ID
	FieldID   int       `json:"field_id"` // 字段 ID
	Value     string    `json:"value"` // 选项值
	Label     string    `json:"label"` // 选项名称
}

type Dependency struct {
	FieldId       string  `json:"field_id"` // 被联动字段的标识
	Value 				string  `json:"value"` // 联动条件值
}
