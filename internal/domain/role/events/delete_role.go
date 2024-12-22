package events

// 角色删除成功事件
type RoleDeletedEvent struct {
	Id string
}

// 角色删除失败事件
type RoleDeleteFailedEvent struct {
	Id string
	Err error
}

