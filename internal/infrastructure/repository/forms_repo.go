package repository

import (
	"encoding/json"

	"github.com/r1005410078/meida-admin-server/internal/app/repository"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/events"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"gorm.io/gorm"
)

type FormsRepository struct {
	db *gorm.DB
}

func NewFormsRepository(db *gorm.DB) repository.IFormsRepository {
	return &FormsRepository{
		db,
	}
}

// SaveForm 保存表单
func (r *FormsRepository) CreateForms(event *events.CreateFormsEvent) error {
	return r.db.Create(&model.Form{
		ID: *event.Id,
		Name: event.Name,
		Description: event.Description,
	}).Error
}

func (r *FormsRepository) UpdateForms(event *events.UpdateFormsEvent) error {
	return r.db.Where("id = ?", event.Id).Updates(&model.Form{
		Name: event.Name,
		Description: event.Description,
	}).Error
}
func (r *FormsRepository) DeleteForms(event *events.DeleteFormsEvent) error {
	return r.db.Where("id = ?", event.Id).Delete(&model.Form{}).Error
}

func (r *FormsRepository) CreateFormsFileds(event *events.CreateFormsFieldsEvent) error {
	newModel, err := ToFormsField(event.SaveFormsFieldsCommand)
	if err != nil {
		return err
	}
	return r.db.Create(&newModel).Error
}

func (r *FormsRepository) UpdateFormsFileds(event *events.UpdateFormsFieldsEvent) error {
	newModel, err := ToFormsField(event.SaveFormsFieldsCommand)
	if err != nil {
		return err
	}
	
	return r.db.Where("form_id = ? and field_id = ?", event.FormID, event.FiledId).
		Updates(&newModel).Error
}

func (r *FormsRepository) DeleteFormsFileds(event *events.DeleteFormsFiledsEvent) error {
	cmd := event.DeleteFormsFieldsCommand
	return r.db.Where("form_id = ? and field_id = ?", cmd.FormID, cmd.FiledId).
		Delete(&model.FormsField{}).Error
}


func (r *FormsRepository) CreateDepends(event *events.CreateDependsOnEvent) error {
	for _, depend := range event.DependsOn {
		dependModel := &model.FormsFieldsDependency{
			ID: *event.Id,
			FormID: event.FormID,
			FieldID: event.FieldID,
			RelatedFieldID: depend.FieldId,
			ConditionValue: depend.Value,
		}

		if err := r.db.Create(dependModel).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *FormsRepository) UpdateDepends(event *events.UpdateDependsOnEvent) error {
	for _, depend := range event.DependsOn {
		dependModel := &model.FormsFieldsDependency{
			FormID: event.FormID,
			FieldID: event.FieldID,
			RelatedFieldID: depend.FieldId,
			ConditionValue: depend.Value,
		}

		if err := r.db.Where("id = ?", event.Id).Updates(dependModel).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *FormsRepository) DeleteDepends(event *events.DeleteDependsOnEvent) error {
	return r.db.Where("id = ?", event.Id).Delete(&model.FormsFieldsDependency{}).Error
}

func ToFormsField(cmd *command.SaveFormsFieldsCommand) (*model.FormsField, error) {
	validationRules, err := json.Marshal(cmd.ValidationRules)
	if err != nil {
		return nil, err
	}

	options, err := json.Marshal(cmd.Options)
	if err != nil {
		return nil, err
	}
	
	validationRulesStr := string(validationRules)
	optionsStr := string(options)
	
	return &model.FormsField{
		FieldID: *cmd.FiledId,
		FormID: *cmd.FormID,
		Label: *cmd.Label,
		Type: *cmd.Type,
		Required: *cmd.Required,
		Placeholder: *cmd.Placeholder,
		ValidationRules: &validationRulesStr,
		Options: &optionsStr,
	}, nil
}