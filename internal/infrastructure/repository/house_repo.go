package repository

import (
	"github.com/r1005410078/meida-admin-server/internal/app/repository"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/events"
	"gorm.io/gorm"
)

type HouseRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewHouseRepository(db *gorm.DB) repository.IHouseRepository {
	return &HouseRepository{
		db: db,
		tx: db,
	}
}

	// 创建房源
func (r *HouseRepository)	CreateHouse(inputHouse *events.CreateHouseEvent) error {
	return nil
}

// 更新房源
func (r *HouseRepository)	UpdateHouse(inputHouse  *events.UpdateHouseEvent) error {
	return nil
}
// 删除房源
func (r *HouseRepository)	DeleteHouse(id *string) error {
	return nil
}

// 创建房源标签
func (r *HouseRepository)	CreateHouseTags(inputHouse *events.CreateHouseTagsEvent) error {
	return nil
}

// 更新房源标签
func (r *HouseRepository)	UpdateHouseTags(inputHouse *events.UpdateHouseTagsEvent) error {
	return nil
}

// 删除房源标签
func (r *HouseRepository)	DeleteHouseTags(inputHouse *events.DeleteHouseTagsEvent) error {
	return nil
}

// 创建房源多媒体
func (r *HouseRepository)	CreateHouseMedias(inputHouse *events.CreateHouseMediasEvent) error {
	return nil
}

// 更新房源多媒体
func (r *HouseRepository)	UpdateHouseMedias(inputHouse *events.UpdateHouseMediasEvent) error {
	return nil
}

// 删除房源多媒体
func (r *HouseRepository)	DeleteHouseMedias(houseId *string) error {
	return nil
}

// 保存房源经纬度
func (r *HouseRepository)	SaveHouseLocation(inputHouse *events.SaveHouseLocationEvent) error {
	return nil
}

// 删除房源经纬度
func (r *HouseRepository)	DeleteHouseLocation(houseId *string) error {
	return nil
}