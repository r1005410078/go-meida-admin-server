package events

import "time"

type LoggedInEvent struct {
	ID        *string    `json:"id"`
	Username  *string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

type LoginFailedEvent struct {
	Username  string   
	Err     error	   
	Timestamp time.Time 
	IP        string    
	LoginAttempts  int32       
}
