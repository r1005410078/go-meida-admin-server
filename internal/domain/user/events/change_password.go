package events

type ChangePasswordEvent struct {
	UserId       string
	PasswordHash string
}

// 修改错误
type ChangePasswordFailedEvent struct {
	ChangePasswordEvent
	Err    error
}

func NewChangePasswordFailedEvent(ChangePasswordEvent ChangePasswordEvent, Err error) *ChangePasswordFailedEvent {
	return &ChangePasswordFailedEvent{
		ChangePasswordEvent,
		Err,
	}
}