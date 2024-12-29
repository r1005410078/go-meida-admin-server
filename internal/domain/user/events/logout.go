package events

type LoggedOutEvent struct {
	UserId string
}

type LoggedOutFailedEvent struct {
	UserId string
	Err    error
}