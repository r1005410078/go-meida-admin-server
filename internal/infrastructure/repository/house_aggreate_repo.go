package repository

import (
	"github.com/r1005410078/meida-admin-server/internal/domain/house"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
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
}

func (r *HouseAggregateRepository) Commit() {
	if r.tx != nil {
		r.tx.Commit()
	} else {
		panic("cannot commit because tx is nil")
	}
	r.tx = nil
}

func (r *HouseAggregateRepository) ExistAddress(address string, id *string) bool {
	var count int64

	if id == nil {
		if err := r.db.Model(&model.HousePropertyAggregate{}).Where("address = ?", address).Count(&count).Error; err != nil {
			return false
		}
	} else {
		if err := r.db.Model(&model.HousePropertyAggregate{}).Where("address = ? and id != ?", address, *id).Count(&count).Error; err != nil {
			return false
		}
	}

	return count > 0
}

func (r *HouseAggregateRepository) GetAggregate(id *string) (*house.HousePropertyAggregate, error) {
	aggregate := model.HousePropertyAggregate{}
	if err := r.db.Where("id = ?", *id).First(&aggregate).Error; err != nil {
		return nil, err
	}

	return aggregate.ToHousePropertyAggregate(), nil
}

func (r *HouseAggregateRepository) SaveAggregate(aggregate *house.HousePropertyAggregate) error {
	db := r.dbInstance()
	res := db.Save(model.FormHousePropertyAggregate(aggregate))

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *HouseAggregateRepository) DeleteAggregate(id *string) error {
	return r.db.Where("id = ?", id).Delete(&model.HousePropertyAggregate{}).Error
}
