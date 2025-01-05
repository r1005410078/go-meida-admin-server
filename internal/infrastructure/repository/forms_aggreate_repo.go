package repository

import (
	"encoding/json"

	"github.com/r1005410078/meida-admin-server/internal/domain/forms"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/command"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"gorm.io/gorm"
)

type FormsAggregateRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewFormsAggregateRepository(db *gorm.DB) forms.IFormsAggregateRepository {
	return &FormsAggregateRepository{
		db: db,
		tx: nil,
	}
}

func (r *FormsAggregateRepository) Begin() {
	r.tx = r.db.Begin()
}

func (r *FormsAggregateRepository) Rollback() {
	if r.tx != nil {
		r.tx.Rollback()
	} else {
		panic("cannot rollback because tx is nil")
	}
}

func (r *FormsAggregateRepository) Commit() {
	if r.tx != nil {
		r.tx.Commit()
	} else {
		panic("cannot commit because tx is nil")
	}
	r.tx = nil
}

func (r *FormsAggregateRepository) dbInstance() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db
}

func (r *FormsAggregateRepository) SaveAggregate(aggregate *forms.FormAggregate) error {
	db := r.dbInstance()

	// 保存字段
	for _, v := range aggregate.Fields {
		if v.CreateAt != nil  {
			if err := db.Create(&model.FormFieldsAggregate{
				FormID: aggregate.FormId,
				FieldID: *v.FieldId,
				Label: *v.FieldName,
			}).Error; err != nil {
				return err
			}
		}

		if v.UpdateAt != nil {
			if err := db.Where("form_id = ? and field_id = ?", aggregate.FormId, *v.FieldId).
				Updates(model.FormFieldsAggregate{
					Label: *v.FieldName,
				}).Error; err != nil {
				return err
			}
		}

		if v.DeleteAt != nil {
			if err := db.Where("form_id = ? and field_id = ?", aggregate.FormId, *v.FieldId).
				Delete(&model.FormFieldsAggregate{}).Error; err != nil {
				return err
			}
		}
	}

	// 保存聚合
	relatedIds, err := json.Marshal(aggregate.RelatedIds)
	if err != nil {
		return err
	}
	relatedIdsStr := string(relatedIds)

	result := db.Where("form_id = ?", aggregate.FormId).Updates(&model.FormsAggregate{
		FormName: aggregate.FormName,
		RelatedIds: &relatedIdsStr,
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return db.Create(&model.FormsAggregate{
			FormID: aggregate.FormId,
			FormName: aggregate.FormName,
		}).Error
	}

	return nil
}


func (r *FormsAggregateRepository) DeleteAggregateFields(cmd *command.DeleteFormsFieldsCommand) error {
	return r.db.Where("form_id = ? and field_id = ?", cmd.FormID, *cmd.FiledId).
		Delete(&model.FormFieldsAggregate{}).Error
}

func (r *FormsAggregateRepository) Updates(aggregate *forms.FormAggregate) error {
	db := r.dbInstance()
	updateDb := db.Model(&model.FormsAggregate{})
	updateDb = updateDb.Where("form_id = ?", aggregate.FormId)
	if err := updateDb.Updates(&aggregate).Error; err != nil {
		return err
	}

	return nil
}

func (r *FormsAggregateRepository) GetAggregate(formId *string) (*forms.FormAggregate, error) {
	var aggregate model.FormsAggregate
	err := r.db.Where("form_id = ?", formId).First(&aggregate).Error
	if err != nil {
		return nil, err
	}

	var formFieldsAggregate []model.FormFieldsAggregate
	if err := r.db.Where("form_id = ?", formId).Find(&formFieldsAggregate).Error; err != nil {
		return nil, err
	}

	fields := []forms.FormField{}
	for _, v := range formFieldsAggregate {
		fields = append(fields, forms.FormField {
			FieldId: &v.FieldID,
			FieldName: &v.Label,
		})
	}

	relatedIdsStr := "[]"
	if aggregate.RelatedIds != nil {
		relatedIdsStr = *aggregate.RelatedIds
	}

	relatedIds := []string{}
	json.Unmarshal([]byte(relatedIdsStr), &relatedIds)

	return &forms.FormAggregate {
		FormId: aggregate.FormID,
		FormName: aggregate.FormName,
		RelatedIds: relatedIds,
		Fields: fields,
	}, nil
}

func (r *FormsAggregateRepository) DeleteAggregateByFormId(formId *string) error {
	err := r.db.Model(&model.FormsAggregate{}).Where("form_id = ?", formId).Delete(&model.FormsAggregate{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *FormsAggregateRepository) DeleteAggregateByFieldId(formId *string) error {
	err := r.db.Model(&model.FormsAggregate{}).Where("form_id = ?", formId).Delete(&model.FormsAggregate{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *FormsAggregateRepository) ExistFieldName(formId, fieldName *string) bool {
	
	var count int64
	r.db.Model(&model.FormsAggregate{}).Where("form_id = ? and field_name = ?", formId, fieldName).Count(&count)
	return count > 0 
}

func (r *FormsAggregateRepository) ExistFormName(formName *string) bool {
	
	var count int64
	r.db.Model(&model.FormsAggregate{}).Where("form_name = ?", formName).Count(&count)
	return count > 0
}

func (r *FormsAggregateRepository) ExistFieldIds(formId string, ids []string) bool {
	var count int64
	r.db.Model(&model.FormsAggregate{}).Where("form_id = ? and field_id in (?)", formId, ids).Count(&count)
	return count > 0
}

