package events

import (
	"github.com/r1005410078/meida-admin-server/internal/domain/house/command"
)

// 保存房源
type CreateHouseEvent struct {
	*command.SaveHouseCommand
}

// 更新房源
type UpdateHouseEvent struct {
	*command.SaveHouseCommand
}

// 删除房源
type DeleteHouseEvent struct {
	*command.DeleteHouseCommand
}

// 添加房源标签
type SaveHouseTagsEvent struct {
	*command.SaveHouseTagsCommand
}

// 添加房源多媒体
type SaveHouseMediasEvent struct {
	*command.SaveHouseMediasCommand
}

// 保存房源经纬度
type SaveHouseLocationEvent struct {
	*command.SaveHouseLocationCommand
}


// 删除房源经纬度
type DeleteHouseLocationEvent struct {
	*command.DeleteHouseLocationCommand
}

// 房源命令错误
type HouseCommandError struct {
	// 命令
  Command	interface {}
	// 错误
	Error error `json:"error"`
}

func NewHouseCommandError(command interface{}, err error) *HouseCommandError {
	return &HouseCommandError{
		Command: command,
		Error:   err,
	}
}