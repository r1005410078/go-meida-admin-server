package events

type UserDeletedEvent struct {
	Id string
}

type UserDeleteFailedEvent struct {
	Id string
	Err error
}