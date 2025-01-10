package house

import (
	"errors"
	"slices"
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
	Address *string
	// 是否在外网同步
	IsSynced *bool
	// tags
	Tags []string
	// 多媒体
	Medias []string
	// 删除时间
	DeletedAt *time.Time
	// 更新时间
	UpdatedAt *time.Time
	// 创建时间
	CreatedAt *time.Time
}

func (h *HousePropertyAggregate) DeleteMedias(command *command.DeleteHouseMediasCommand) (any, error) {
	if !slices.Contains(h.Medias, *command.ID) {
		return nil, errors.New("房源多媒体资源不存在")
	}

	return events.DeleteHouseMediasEvent{DeleteHouseMediasCommand: command}, nil
}

func (h *HousePropertyAggregate) SaveMedias(command *command.SaveHouseMediasCommand) (any, error) {
	if command.ID == nil {
		newMedia := shared.NewId()
		h.Medias = append(h.Medias, *newMedia)
		command.ID = newMedia
		return events.CreateHouseMediasEvent{SaveHouseMediasCommand: command}, nil
	} else {
		if !slices.Contains(h.Medias, *command.ID) {
			return nil, errors.New("房源多媒体资源不存在")
		}
	}

	return events.UpdateHouseMediasEvent{SaveHouseMediasCommand: command}, nil
}

func (h *HousePropertyAggregate) DeleteTag(command *command.DeleteHouseTagsCommand) (any, error) {
	if !slices.Contains(h.Tags, *command.ID) {
		return nil, errors.New("房源标签不存在")
	}

	return events.DeleteHouseTagsEvent{DeleteHouseTagsCommand: command}, nil
}

func (h *HousePropertyAggregate) SaveTags(command *command.SaveHouseTagsCommand) (any, error) {
	if command.ID == nil {
		newTag := shared.NewId()
		h.Tags = append(h.Tags, *newTag)
		command.ID = newTag
		return events.CreateHouseTagsEvent{SaveHouseTagsCommand: command}, nil
	} else {
		if !slices.Contains(h.Tags, *command.ID) {
			return nil, errors.New("id 不存在")
		}
		return events.UpdateHouseTagsEvent{SaveHouseTagsCommand: command}, nil
	}
}

func (h *HousePropertyAggregate) Update(command *command.SaveHouseCommand) any {
	now := time.Now()
	h.Address = command.ToAddress()
	h.IsSynced = command.HouseDetails.ExternalSync
	h.UpdatedAt = &now

	return events.UpdateHouseEvent{SaveHouseCommand: command}
}

func New(cmd *command.SaveHouseCommand) *HousePropertyAggregate {
	now := time.Now()
	return &HousePropertyAggregate{
		HousePropertyId: shared.NewId(),
		Address:         cmd.ToAddress(),
		IsSynced:        cmd.HouseDetails.ExternalSync,
		CreatedAt:       &now,
	}
}

// 删除聚合
func (h *HousePropertyAggregate) Delete(command *command.DeleteHouseCommand) any {
	now := time.Now()
	h.DeletedAt = &now
	return events.DeleteHouseEvent{}
}
