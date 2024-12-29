package events

import "time"

type LoggedInEvent struct {
	ID        *string    `json:"id"`
	Username  *string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

type LoginFailedEvent struct {
	Username  string    `json:"username"`
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	Attempts  int32       `json:"attempts"`
}
