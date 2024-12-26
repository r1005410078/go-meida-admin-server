package events

type UserStatusEvent struct {
	Id string
	Status string
}

type UserStatusFailedEvent struct {
	UserStatusEvent
	Err error
}

func NewUserStatusFailedEvent(UserStatusEvent *UserStatusEvent, Err error) UserStatusFailedEvent {
	return UserStatusFailedEvent{
		*UserStatusEvent,
		Err,
	}
}