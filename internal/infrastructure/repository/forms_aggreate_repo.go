package repository

import (
	"github.com/r1005410078/meida-admin-server/internal/domain/forms"
	"gorm.io/gorm"
)

type FormsAggregateRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewFormsAggregateRepository(db *gorm.DB) *FormsAggregateRepository {
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
}

func (r *FormsAggregateRepository) dbInstance() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db
}

func (r *FormsAggregateRepository) Save(aggregate *forms.FormAggregate) error {
	// Implement the Save method here
	// You can use the r.dbInstance() to get the database instance
	// and perform the necessary operations to save the aggregate.
	return nil
}

func (r *FormsAggregateRepository) Updates(aggregate []*forms.FormAggregate) error {
	// Implement the Updates method here
	// You can use the r.dbInstance() to get the database instance
	// and perform the necessary operations to update the aggregates.
	return nil
}

func (r *FormsAggregateRepository) GetAggregate(formId, fieldId *string) (*forms.FormAggregate, error) {
	// Implement the GetAggregate method here
	// You can use the r.dbInstance() to get the database instance
	// and perform the necessary operations to retrieve the aggregate.
	return nil, nil
}

func (r *FormsAggregateRepository) GetAggregates(formId *string) ([]*forms.FormAggregate, error) {
	// Implement the GetAggregates method here
	// You can use the r.dbInstance() to get the database instance
	// and perform the necessary operations to retrieve the aggregates.
	return nil, nil
}

func (r *FormsAggregateRepository) DeleteAggregate(id *string) error {
	// Implement the DeleteAggregate method here
	// You can use the r.dbInstance() to get the database instance
	// and perform the necessary operations to delete the aggregate.
	return nil
}

func (r *FormsAggregateRepository) ExistFieldName(formId, fieldName *string) bool {
	// Implement the ExistFieldName method here
	// You can use the r.dbInstance() to get the database instance
	// and perform the necessary operations to check if the field name exists.
	return false
}

func (r *FormsAggregateRepository) ExistFormName(formName *string) bool {
	// Implement the ExistFormName method here
	// You can use the r.dbInstance() to get the database instance
	// and perform the necessary operations to check if the form name exists.
	return false
}

func (r *FormsAggregateRepository) ExistFieldIds(formId string, ids []string) bool {
	// Implement the ExistFieldIds method here
	// You can use the r.dbInstance() to get the database instance
	// and perform the necessary operations to check if the field ids exist.
	return false
}