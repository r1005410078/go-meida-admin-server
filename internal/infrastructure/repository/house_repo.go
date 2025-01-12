package repository

import (
	"github.com/r1005410078/meida-admin-server/internal/app/repository"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/events"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
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
	house := model.FormCreateHouseEvent(inputHouse)
	return r.db.Create(house).Error
}
 
// 更新房源
func (r *HouseRepository)	UpdateHouse(inputHouse  *events.UpdateHouseEvent) error {
	r.db.Updates(model.FormUpdateHouseEvent(inputHouse))
	return nil
}

// 删除房源
func (r *HouseRepository)	DeleteHouse(id *string) error {
	return r.db.Where("id = ?", id).Delete(&model.HouseProperty{}).Error
}

// 保存房源标签
func (r *HouseRepository)	SaveHouseTags(inputHouse *events.SaveHouseTagsEvent) error {
  tx :=	r.db.Begin()
	defer tx.Commit()

	// 删除
	if err := r.db.Where("house_property_id = ?", *inputHouse.HousePropertyID).Delete(&model.HousePropertyTag{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 保存
	tags := []model.HousePropertyTag{}
	for _, tag := range inputHouse.Tags {
		tags = append(tags, model.HousePropertyTag{
			HousePropertyID: *inputHouse.HousePropertyID,
			Tag: tag,
		})
	}

	return r.db.Create(tags).Error
}

// 保存房源多媒体
func (r *HouseRepository)	SaveHouseMedias(inputHouse *events.SaveHouseMediasEvent) error {
	tx :=	r.db.Begin()
	defer tx.Commit()

	// 删除
	if err := r.db.Where("house_property_id = ?", inputHouse.HousePropertyID).Delete(&model.HousePropertyMedia{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 保存
	urls := []model.HousePropertyMedia{}
	for _, url := range inputHouse.Urls {
		urls = append(urls, model.HousePropertyMedia{
			HousePropertyID: *inputHouse.HousePropertyID,
			URL: url,
		})
	}

	return r.db.Create(urls).Error
}

// 保存房源经纬度
func (r *HouseRepository)	SaveHouseLocation(inputHouse *events.SaveHouseLocationEvent) error {
	result := r.db.Where("house_property_id = ?", *inputHouse.HousePropertyID).Updates(&model.HousePropertyLocation {
		HousePropertyID: *inputHouse.HousePropertyID,
		Latitude: inputHouse.Latitude,
		Longitude: inputHouse.Longitude,
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return r.db.Create(&model.HousePropertyLocation {
			HousePropertyID: *inputHouse.HousePropertyID,
			Latitude: inputHouse.Latitude,
			Longitude: inputHouse.Longitude,
		}).Error
	}

	return nil
}

// 删除房源经纬度
func (r *HouseRepository)	DeleteHouseLocation(houseId *string) error {
	return  r.db.Where("house_property_id = ?", *houseId).Delete(&model.HousePropertyLocation{}).Error
}