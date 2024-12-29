package events

import "time"

type RegisteredEvent struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterFailedEvent struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}
