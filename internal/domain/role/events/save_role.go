package events

type RoleSavedEvent struct {
	// 角色ID
	Id *string
	// 角色名称
	Name string
	// 角色权限
	PermissionIds []string
	// 角色描述
	Description *string
}

// 角色保存失败事件
type RoleSaveFailedEvent struct {
	RoleSavedEvent
	// 错误信息
	Err error
}

func NewRoleSaveFailedEvent(RoleSavedEvent RoleSavedEvent, err error) RoleSaveFailedEvent {
	return RoleSaveFailedEvent{
		RoleSavedEvent: RoleSavedEvent,
		Err: err,
	}
}