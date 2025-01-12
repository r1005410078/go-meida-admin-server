package values

import "time"


type Purpose string

type HouseDetails struct {
	// 楼层
	FloorNumber *int32 `json:"floorNumber"`
	// 起始楼层
	FloorNumberFrom *int32 `json:"floorNumberFrom"`
	// 结束楼层
	FloorNumberTo *int32 `json:"floorNumberTo"`
	// 房源标题
	Title *string `json:"title"`
	// 车位高度
	CarHeight *float64 `json:"carHeight"`
	// 户型-房
	LayoutRoom *int32 `json:"layoutRoom"`
	// 户型-厅
	LayoutHall *int32 `json:"layoutHall"`
	// 户型-餐
	LayoutKitchen *int32 `json:"layoutKitchen"`
	// 户型-卫
	LayoutBathroom *int32 `json:"layoutBathroom"`
	// 户型-阳台
	LayoutBalcony *int32 `json:"layoutBalcony"`
	// 梯
	Stairs *int32 `json:"stairs"`
	// 户
	Rooms *int32 `json:"rooms"`
	// 实率
	ActualRate *float64 `json:"actualRate"`
	// 级别
	Level *int32 `json:"level"`
	// 层高
	FloorHeight *float64 `json:"floorHeight"`
	// 进深
	ProgressDepth *float64 `json:"progressDepth"`
	// 门宽
	DoorWidth *float64 `json:"doorWidth"`
	// 建筑面积
	BuildingArea int32 `json:"buildingArea"`
	// 使用面积
	UseArea float64 `json:"useArea"`
	// 售价
	SalePrice float64 `json:"salePrice"`
	// 租价
	RentPrice float64 `json:"rentPrice"`
	// 出租低价
	RentLowPrice *float64 `json:"rentLowPrice"`
	// 首付
	DownPayment *float64 `json:"downPayment"`
	// 出售低价
	SaleLowPrice *float64 `json:"saleLowPrice"`
	// 房屋类型
	HouseType *string `json:"houseType"`
	// 房屋朝向
	HouseOrientation *string `json:"houseOrientation"`
	// 装修
	HouseDecoration *string `json:"houseDecoration"`
	// 满减年限
	DiscountYearLimit *int32 `json:"discountYearLimit"`
	// 看房方式
	ViewMethod *string `json:"viewMethod"`
	// 付款方式
	PaymentMethod *string `json:"paymentMethod"`
	// 房源税费
	PropertyTax *float64 `json:"propertyTax"`
	// 建筑结构
	BuildingStructure *string `json:"buildingStructure"`
	// 建筑年代
	BuildingYear *string `json:"buildingYear"`
	// 产权性质
	PropertyRights *string `json:"propertyRights"`
	// 产权年限
	PropertyYearLimit *int32 `json:"propertyYearLimit"`
	// 产权日期
	CertificateDate *time.Time `json:"certificateDate"`
	// 交房日期
	HandoverDate *time.Time `json:"handoverDate"`
	// 学位
	Degree *string `json:"degree"`
	// 户口
	Household *string `json:"household"`
	// 来源
	Source *string `json:"source"`
	// 委托编号
	DelegateNumber *string `json:"delegateNumber"`
	// 唯一住房
	UniqueHousing *bool `json:"uniqueHousing"`
	// 全款
	FullPayment *bool `json:"fullPayment"`
	// 抵押
	Mortgage *bool `json:"mortgage"`
	// 急切
	Urgent *bool `json:"urgent"`
	// 配套
	Support *string `json:"support"`
	// 现状
	PresentState *string `json:"presentState"`
	// 外网同步
	ExternalSync bool `json:"externalSync"`
	// 备注
	Remark *string `json:"remark"`
	// 房源标签
	Tags *[]string `json:"tags"`
	// 房源
	Medias *[]string `json:"medias"`
	// 房源经纬度
	*Location
}

type Location struct {
	Latitude *float64 `json:"latitude" binding:"required"`
	Longitude *float64 `json:"longitude" binding:"required"`
}