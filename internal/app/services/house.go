package services

import (
	"github.com/r1005410078/meida-admin-server/internal/app/repository"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/events"
	"go.uber.org/zap"
)

type HouseServices struct {
	repo repository.IHouseRepository
	logger *zap.Logger
}

func NewHouseServices(repo repository.IHouseRepository, logger *zap.Logger) *HouseServices {
	return &HouseServices{
		repo: repo,
		logger: logger,
	}
}


// 获取房源列表
// func (s *HouseServices) ListHouse() ([]*model., error) {
// 	return s.repo.ListHouse(event)
// }

// 保存房源
func (s *HouseServices) CreateHouseEventHandle(event *events.CreateHouseEvent) error {
	return s.repo.CreateHouse(event)
}

// 更新房源
func (s *HouseServices) UpdateHouseEventHandle(event *events.UpdateHouseEvent) error {
	return s.repo.UpdateHouse(event)
}

// 删除房源
func (s *HouseServices) DeleteHouseEventHandle(event *events.DeleteHouseEvent) error {
	return s.repo.DeleteHouse(event.ID)
}

// 保存房源标签
func (s *HouseServices) SaveHouseTagsEventHandle(event *events.SaveHouseTagsEvent) error {
	return s.repo.SaveHouseTags(event)
}

// 保存房源多媒体
func (s *HouseServices) SaveHouseMediasEventHandle(event *events.SaveHouseMediasEvent) error {
	return  s.repo.SaveHouseMedias(event)
}

// 保存房源经纬度
func (s *HouseServices) SaveHouseLocationEventHandle(event *events.SaveHouseLocationEvent) error {
	return s.repo.SaveHouseLocation(event)
}

// 删除房源经纬度
func (s *HouseServices) DeleteHouseLocationEventHandle(event *events.DeleteHouseLocationEvent) error {
	return s.repo.DeleteHouseLocation(event.ID)
}

// 房源命令错误
func (s *HouseServices) HouseCommandErrorEventHandle(event *events.HouseCommandError) error {
	s.logger.Error("house command error", zap.Any("event", event))
	return nil
}