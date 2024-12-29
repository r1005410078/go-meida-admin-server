package events

type LoggedOutEvent struct {
	UserId string
}

type LoggedOutFailedEvent struct {
	Token string
	UserId string
	Err    error
}