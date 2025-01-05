package events

import "github.com/r1005410078/meida-admin-server/internal/domain/forms/command"


type CreateDependsOnEvent struct {
	*command.SaveDependsOnCommand
}

type UpdateDependsOnEvent struct {
	*command.SaveDependsOnCommand
}

type DeleteDependsOnEvent struct {
	Id 						   string
}
