package repository

import "github.com/r1005410078/meida-admin-server/internal/domain/house/events"

type IHouseRepository interface {
	// 创建房源
	CreateHouse(inputHouse *events.CreateHouseEvent) error
	// 更新房源
	UpdateHouse(inputHouse  *events.UpdateHouseEvent) error
	// 删除房源
	DeleteHouse(id *string) error

	// 创建房源标签
	CreateHouseTags(inputHouse *events.CreateHouseTagsEvent) error
	// 更新房源标签
	UpdateHouseTags(inputHouse *events.UpdateHouseTagsEvent) error
	// 删除房源标签
	DeleteHouseTags(inputHouse *events.DeleteHouseTagsEvent) error

	// 创建房源多媒体
	CreateHouseMedias(inputHouse *events.CreateHouseMediasEvent) error
	// 更新房源多媒体
	UpdateHouseMedias(inputHouse *events.UpdateHouseMediasEvent) error
	// 删除房源多媒体
	DeleteHouseMedias(houseId *string) error

	// 保存房源经纬度
	SaveHouseLocation(inputHouse *events.SaveHouseLocationEvent) error
	// 删除房源经纬度
	DeleteHouseLocation(houseId *string) error
}