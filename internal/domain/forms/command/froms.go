package command

// 保存表单
type SaveFormsCommand struct {
	Id *string `json:"id"`
	Name string `json:"name" binding:"required"`
	Description *string `json:"description"`
}

// 删除表单
type DeleteFormsCommand struct {
	Id string `json:"id" binding:"required"`
}