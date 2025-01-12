package model

import (
	"github.com/r1005410078/meida-admin-server/internal/domain/house"
)

func FormHousePropertyAggregate(e *house.HousePropertyAggregate) HousePropertyAggregate {
	return HousePropertyAggregate{
		ID:        *e.HousePropertyId,
		Address:   e.Address,
		IsSynced:  e.IsSynced,
	}
}

func (h *HousePropertyAggregate) ToHousePropertyAggregate() *house.HousePropertyAggregate {
	return &house.HousePropertyAggregate{
		HousePropertyId: &h.ID,
		Address:   h.Address,
		IsSynced:  h.IsSynced,
	}
}