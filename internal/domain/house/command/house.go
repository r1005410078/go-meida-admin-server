package command

import (
	"fmt"

	"github.com/r1005410078/meida-admin-server/internal/domain/house/values"
)

// 保存房源
type SaveHouseCommand struct {
	ID                 *string `json:"id"`
	// 用途
  Purpose            *string 	`json:"purpose"`
	// 交易类型
	TransactionType    *string 	`json:"transactionType"`
	// 状态
	HouseStatus        *string 	`json:"houseStatus"`
	// 业主姓名
	OwnerName          *string 	`json:"ownerName"`
	// 联系电话
	Phone              *string 	`json:"phone"`
	// 小区
	Community          *string 	`json:"community" binding:"required"`
	// 起始楼层
	FloorRangeMin      *int 	`json:"floorRangeMin"`
	// 结束楼层
	FloorRangeMax      *int 	`json:"floorRangeMax"`
  // 座栋
	BuildingNumber     *int 	`json:"building_number" binding:"required"`
	// 单元
	UnitNumber         *int 	`json:"unit_number" binding:"required"`
	// 门牌号  
	DoorNumber         *int 	`json:"door_number" binding:"required"`
	// 房源详情
	HouseDetails       *values.HouseDetails `json:"house_details"`
}

func (c *SaveHouseCommand) ToAddress() *string {
	address := *c.Community + "," + fmt.Sprint(*c.BuildingNumber) + "-" + fmt.Sprint(*c.UnitNumber) + "-" + fmt.Sprint(*c.DoorNumber)
	return &address
}

// 删除房源
type DeleteHouseCommand struct {
	ID *string `json:"id" binding:"required"`
}

// 添加房源标签
type SaveHouseTagsCommand struct {
	ID *string `json:"id"`
	Name *string `json:"name" binding:"required"`
}

// 删除房源标签
type DeleteHouseTagsCommand struct {
	ID *string `json:"id"  binding:"required"`
}

// 保存房源多媒体
type SaveHouseMediasCommand struct {
	ID *string `json:"id"`
	Url *string `json:"url"  binding:"required"`
}

// 删除房源多媒体
type DeleteHouseMediasCommand struct {
	ID *string `json:"id"  binding:"required"`
}

// 保存房源经纬度
type SaveHouseLocationCommand struct {
	ID *string `json:"id"`
	Latitude *float64 `json:"latitude" binding:"required"`
	Longitude *float64 `json:"longitude" binding:"required"`
}

// 删除房源经纬度
type DeleteHouseLocationCommand struct {
	ID *string `json:"id" binding:"required"`
}