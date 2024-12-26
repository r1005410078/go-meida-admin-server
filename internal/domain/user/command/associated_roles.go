package command

import "github.com/r1005410078/meida-admin-server/internal/domain/user/events"

type AssociatedRolesCommand struct {
	UserId string  `json:"userId" binding:"required"`
	RoleId string  `json:"roleId" binding:"required"`
}
 
func (command *AssociatedRolesCommand) ToEvent() *events.AssoicatedRolesEvent {
	return &events.AssoicatedRolesEvent{
		UserId: command.UserId,
		RoleId: command.RoleId,
	}
}