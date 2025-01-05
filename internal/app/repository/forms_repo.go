package repository

import "github.com/r1005410078/meida-admin-server/internal/domain/forms/events"

type IFormsRepository interface {
	CreateForms(event *events.CreateFormsEvent) error
	UpdateForms(event *events.UpdateFormsEvent) error
	DeleteForms(event *events.DeleteFormsEvent) error
	CreateFormsFileds(event *events.CreateFormsFieldsEvent) error
	UpdateFormsFileds(event *events.UpdateFormsFieldsEvent) error
	DeleteFormsFileds(event *events.DeleteFormsFiledsEvent) error

	CreateDepends(event *events.CreateDependsOnEvent) error
	UpdateDepends(event *events.UpdateDependsOnEvent) error
	DeleteDepends(event *events.DeleteDependsOnEvent) error
}