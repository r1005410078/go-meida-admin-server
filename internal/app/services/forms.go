package services

import (
	"github.com/r1005410078/meida-admin-server/internal/app/repository"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
	"go.uber.org/zap"
)

type FormsServices struct {
	repo repository.IFormsRepository
	aggRepo forms.IFormsAggregateRepository
	eventBus shared.IEventBus
	logger *zap.Logger
}

func NewFormsServices(repo repository.IFormsRepository, logger *zap.Logger) *FormsServices {
	return &FormsServices{
		repo: repo,
		logger: logger,
	}
}

func (s *FormsServices) CreateFormsEventHandle(event *events.CreateFormsEvent) error {
	return s.repo.CreateForms(event)
}

func (s *FormsServices) UpdateFormsEventHandle(event *events.UpdateFormsEvent) error {
	return s.repo.UpdateForms(event)
}

func (s *FormsServices) SaveFormsFailedEventHandle(event *events.SaveFormsFailedEvent) error {
	s.logger.Error("save forms failed", zap.Error(event.Err))
	return event.Err
}

func (s *FormsServices) DeleteFormsEventHandle(event *events.DeleteFormsEvent) error {
	return s.repo.DeleteForms(event)
}

func (s *FormsServices) DeleteFormsFailedEventHandle(event *events.DeleteFormsFailedEvent) error {
	s.logger.Error("delete forms failed", zap.Error(event.Err))
	return event.Err
}

/// forms fileds
func (s *FormsServices) CreateFormsFieldsEventHanlde(event *events.CreateFormsFieldsEvent) error {
	return s.repo.CreateFormsFileds(event)
}

func (s *FormsServices) UpdateFormsFieldsEventHanlde(event *events.UpdateFormsFieldsEvent) error {
	return s.repo.UpdateFormsFileds(event)
}

func (s *FormsServices) DeleteFormsFiledsEventHanlde(event *events.DeleteFormsFiledsEvent) error {
	return s.repo.DeleteFormsFileds(event)
}

// error
func (s *FormsServices) CreateFormsFiledsFailedEventHanlde(event *events.CreateFormsFiledsFailedEvent) error {
	s.logger.Error("create forms fileds failed", zap.Error(event.Err))
	return event.Err
}

func (s *FormsServices) UpdateFormsFieldsFailedEventHanlde(event *events.UpdateFormsFieldsFailedEvent) error {
	s.logger.Error("update forms fileds failed", zap.Error(event.Err))
	return event.Err
}

func (s *FormsServices) DeleteFormsFiledsFailedEventHanlde(event *events.DeleteFormsFiledsFailedEvent) error {
	s.logger.Error("delete forms fileds failed", zap.Error(event.Err))
	return event.Err
}

/// depends on
func (s *FormsServices) CreateDependsOnEventHandle(event *events.CreateDependsOnEvent) error {
	return s.repo.CreateDepends(event)
}

func (s *FormsServices) UpdateDependsOnEventHandle(event *events.UpdateDependsOnEvent) error {
	return s.repo.UpdateDepends(event)
}

func (s *FormsServices) DeleteDependsOnEventHandle(event *events.DeleteDependsOnEvent) error {
	return s.repo.DeleteDepends(event)
}