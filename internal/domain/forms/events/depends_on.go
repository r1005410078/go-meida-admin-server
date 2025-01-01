package events

import "github.com/r1005410078/meida-admin-server/internal/domain/forms/command"


type SaveDependsOnEvent struct {
	*command.SaveDependsOnCommand
}

type ApeendDependsOnEvent struct {
	*command.SaveDependsOnCommand
}

type DeleteDependsOnEvent struct {
	*command.SaveDependsOnCommand
}
