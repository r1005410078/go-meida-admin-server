package house

import (
	"errors"
	"time"

	"github.com/r1005410078/meida-admin-server/internal/domain/house/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/events"
	"github.com/r1005410078/meida-admin-server/internal/domain/shared"
)

// 房源聚合
type HousePropertyAggregate struct {
	// 房源ID
	HousePropertyId *string
	// 地址
	Address string
	// 是否在外网同步
	IsSynced bool
	// 删除时间
	DeletedAt *time.Time
	// 更新时间
	UpdatedAt *time.Time
	// 创建时间
	CreatedAt *time.Time
}


func (h *HousePropertyAggregate) Update(command *command.SaveHouseCommand) (any, error) {
	now := time.Now()
	h.Address = command.ToAddress()
	if command.HouseDetails != nil {
		h.IsSynced = command.HouseDetails.ExternalSync
	}
	h.UpdatedAt = &now

	if err := validateFloorRange(command); err != nil {
		return nil, err
	}

	return &events.UpdateHouseEvent{SaveHouseCommand: command}, nil
}

func New(cmd *command.SaveHouseCommand) (*HousePropertyAggregate, error) {
	now := time.Now()
	house := &HousePropertyAggregate{
		HousePropertyId: shared.NewId(),
		Address:         cmd.ToAddress(),
		CreatedAt:       &now,
	}

	if cmd.HouseDetails != nil {
		house.IsSynced = cmd.HouseDetails.ExternalSync
	}

	if err := validateFloorRange(cmd); err != nil {
		return nil, err
	}

	return house, nil
}

func validateFloorRange(command *command.SaveHouseCommand) error {
	if command.FloorRangeMin != nil && command.FloorRangeMax != nil {
		if *command.FloorRangeMin >= *command.FloorRangeMax {
			return errors.New("楼层范围错误")
		}
	}
	return nil
}

// 删除聚合
func (h *HousePropertyAggregate) Delete(command *command.DeleteHouseCommand) any {
	now := time.Now()
	h.DeletedAt = &now
	return &events.DeleteHouseEvent{
		DeleteHouseCommand: command,
	}
}
