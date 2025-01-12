package command

import (
	"fmt"

	"github.com/r1005410078/meida-admin-server/internal/domain/house/values"
)

// 保存房源
type SaveHouseCommand struct {
	ID                 *string  `json:"id"`
	// 用途
  Purpose            *string 	`json:"purpose"`
	// 交易类型
	TransactionType    *string 	`json:"transactionType"`
	// 状态
	HouseStatus        *string 	`json:"houseStatus"`
	// 业主姓名
	OwnerName          string 	`json:"ownerName"`
	// 联系电话
	Phone              string 	`json:"phone"`
	// 小区
	Community          string 	`json:"community"`
	// 起始楼层
	FloorRangeMin      *int32 	`json:"floorRangeMin"`
	// 结束楼层
	FloorRangeMax      *int32 	`json:"floorRangeMax"`
  // 座栋
	BuildingNumber     *int32 	`json:"buildingNumber"`
	// 单元
	UnitNumber         *int32 	`json:"unitNumber"`
	// 门牌号  
	DoorNumber         *int32 	`json:"doorNumber"`
	// 房源详情
	HouseDetails       *values.HouseDetails `json:"houseDetails"`
}

func (c *SaveHouseCommand) ToAddress() string {
	address := c.Community + "," + fmt.Sprint(*c.BuildingNumber) + "-" + fmt.Sprint(*c.UnitNumber) + "-" + fmt.Sprint(*c.DoorNumber)
	return address
}

func (c *SaveHouseCommand) ToSaveHouseTagsCommand() *SaveHouseTagsCommand {
	return &SaveHouseTagsCommand {
		HousePropertyID: c.ID,
		Tags: *c.HouseDetails.Tags,
	}
}

func (c *SaveHouseCommand) ToSaveHouseMediasCommand() *SaveHouseMediasCommand {
	return &SaveHouseMediasCommand {
		HousePropertyID: c.ID,
		Urls: *c.HouseDetails.Medias,
	}
}

func (c *SaveHouseCommand) ToSaveHouseLocationCommand() *SaveHouseLocationCommand {
	return &SaveHouseLocationCommand {
		HousePropertyID: c.ID,
		Latitude: c.HouseDetails.Latitude,
		Longitude: c.HouseDetails.Longitude,
	}
}

// 删除房源
type DeleteHouseCommand struct {
	ID *string `json:"id" binding:"required"`
}

// 保存房源标签
type SaveHouseTagsCommand struct {
	HousePropertyID *string `json:"houseId" binding:"required"`
	Tags []string `json:"tags" binding:"required"`
}

// 保存房源多媒体
type SaveHouseMediasCommand struct {
	HousePropertyID *string `json:"houseId"`
	Urls []string `json:"urls"  binding:"required"`
}

// 保存房源经纬度
type SaveHouseLocationCommand struct {
	HousePropertyID *string `json:"houseId"`
	Latitude *float64 `json:"latitude" binding:"required"`
	Longitude *float64 `json:"longitude" binding:"required"`
}

// 删除房源经纬度
type DeleteHouseLocationCommand struct {
	ID *string `json:"id" binding:"required"`
}