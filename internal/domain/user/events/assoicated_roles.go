package events

type AssoicatedRolesEvent struct {
	UserId string 
	RoleId string 
}

type AssoicatedRolesFailedEvent struct {
	*AssoicatedRolesEvent
	Err error 
} 

func NewAssoicatedRolesFailedEvent(AssoicatedRolesEvent *AssoicatedRolesEvent, Err error) *AssoicatedRolesFailedEvent {
	return &AssoicatedRolesFailedEvent{
		AssoicatedRolesEvent,
		Err,
	}
}