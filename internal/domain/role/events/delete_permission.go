package events

// 角色删除成功事件
type DeletedPermissionEvent struct {
	// 角色id
	Id string
	// 权限id
	PermissionId []string
}

// 角色删除失败事件
type DeletePermissionFailedEvent struct {
	// 角色id
	Id string
	// 权限id
	PermissionId []string
	// 错误信息
	Err error
}

func NewDeletePermissionFailedEvent (Id string, PermissionId []string, Err error ) DeletePermissionFailedEvent {
	return DeletePermissionFailedEvent{
		Id,
		PermissionId,
		Err,
	}
}