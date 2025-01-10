package values


type Purpose string

type HouseDetails struct {
	// 楼层
	FloorNumber *int `json:"floor_number"`
	// 起始楼层
	FloorNumberFrom *int `json:"floor_number_from"`
	// 结束楼层
	FloorNumberTo *int `json:"floor_number_to"`
	// 房源标题
	Title *string `json:"title"`
	// 车位高度
	CarHeight *float64 `json:"car_height"`
	// 户型-房
	LayoutRoom *int `json:"layout_room"`
	// 户型-厅
	LayoutHall *int `json:"layout_hall"`
	// 户型-餐
	LayoutKitchen *int `json:"layout_kitchen"`
	// 户型-卫
	LayoutBathroom *int `json:"layout_bathroom"`
	// 户型-阳台
	LayoutBalcony *int `json:"layout_balcony"`
	// 梯
	Stairs *int `json:"stairs"`
	// 户
	Rooms *int `json:"rooms"`
	// 实率
	ActualRate *float64 `json:"actual_rate"`
	// 级别
	Level *int `json:"level"`
	// 层高
	FloorHeight *float64 `json:"floor_height"`
	// 进深
	ProgressDepth *float64 `json:"progress_depth"`
	// 门宽
	DoorWidth *float64 `json:"door_width"`
	// 建筑面积
	BuildingArea *int `json:"building_area"`
	// 使用面积
	UseArea *float64 `json:"use_area"`
	// 售价
	SalePrice *float64 `json:"sale_price"`
	// 租价
	RentPrice *float64 `json:"rent_price"`
	// 出租低价
	RentLowPrice *float64 `json:"rent_low_price"`
	// 首付
	DownPayment *float64 `json:"down_payment"`
	// 出售低价
	SaleLowPrice *float64 `json:"sale_low_price"`
	// 房屋类型
	HouseType *string `json:"house_type"`
	// 房屋朝向
	HouseOrientation *string `json:"house_orientation"`
	// 装修
	HouseDecoration *string `json:"house_decoration"`
	// 满减年限
	DiscountYearLimit *int `json:"discount_year_limit"`
	// 看房方式
	ViewMethod *string `json:"view_method"`
	// 付款方式
	PaymentMethod *string `json:"payment_method"`
	// 房源税费
	PropertyTax *float64 `json:"property_tax"`
	// 建筑结构
	BuildingStructure *string `json:"building_structure"`
	// 建筑年代
	BuildingYear *string `json:"building_year"`
	// 产权性质
	PropertyRights *string `json:"property_rights"`
	// 产权年限
	PropertyYearLimit *int `json:"property_year_limit"`
	// 产权日期
	CertificateDate *string `json:"certificate_date"`
	// 交房日期
	HandoverDate *string `json:"handover_date"`
	// 学位
	Degree *string `json:"degree"`
	// 户口
	Household *string `json:"household"`
	// 来源
	Source *string `json:"source"`
	// 委托编号
	DelegateNumber *string `json:"delegate_number"`
	// 唯一住房
	UniqueHousing *bool `json:"unique_housing"`
	// 全款
	FullPayment *bool `json:"full_payment"`
	// 抵押
	Mortgage *bool `json:"mortgage"`
	// 急切
	Urgent *bool `json:"urgent"`
	// 配套
	Support *string `json:"support"`
	// 现状
	PresentState *string `json:"present_state"`
	// 外网同步
	ExternalSync *bool `json:"external_sync"`
	// 备注
	Remark *string `json:"remark"`
	// 房源标签
	Tags *[]string `json:"tags"`
}