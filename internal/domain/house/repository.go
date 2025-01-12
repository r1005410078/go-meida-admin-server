package house

type IHouseAggregateRepository interface {

	Begin()
	Rollback()
	Commit()

	// 地址是否存在
	ExistAddress(address string, id *string) bool

	// 聚合
	GetAggregate(id *string) (*HousePropertyAggregate, error)
	SaveAggregate(aggregate *HousePropertyAggregate) error
	DeleteAggregate(id *string) error
}