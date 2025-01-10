package repository

import (
	"github.com/r1005410078/meida-admin-server/internal/domain/house"
	"gorm.io/gorm"
)

type HouseAggregateRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewHouseAggregateRepository(db *gorm.DB) house.IHouseAggregateRepository {
	return &HouseAggregateRepository{
		db: db,
		tx: nil,
	}
}

func (r *HouseAggregateRepository) dbInstance() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db
}

func (r *HouseAggregateRepository) Begin() {
	r.tx = r.db.Begin()
}

func (r *HouseAggregateRepository) Rollback() {
	if r.tx != nil {
		r.tx.Rollback()
	} else {
		panic("cannot rollback because tx is nil")
	}
	r.tx = nil
}

func (r *HouseAggregateRepository) Commit() {
	if r.tx != nil {
		r.tx.Commit()
	} else {
		panic("cannot commit because tx is nil")
	}
	r.tx = nil
}

func (r *HouseAggregateRepository) ExistAddress(address *string) bool {
	return false
}

func (r *HouseAggregateRepository) GetAggregate(id *string) (*house.HousePropertyAggregate, error) {
	return nil, nil
}

func (r *HouseAggregateRepository) SaveAggregate(aggregate *house.HousePropertyAggregate) error {
	return nil
}

func (r *HouseAggregateRepository) DeleteAggregate(id *string) error {
	return nil
}
