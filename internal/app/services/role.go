package services

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/app/repository"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/events"
	"go.uber.org/zap"
)

type RoleServices struct {
	repo repository.IRoleRepository
	logger *zap.Logger
}

func NewRepoServices(repo repository.IRoleRepository, logger *zap.Logger) *RoleServices {
	return &RoleServices{
		repo,
		logger,
	}
}

// 保存角色事件处理
func (s *RoleServices) SaveRoleEventHandle(event events.RoleSavedEvent) error {
	if err := s.repo.SaveRole(event); err != nil {
		return errors.New("save role failed " + err.Error())
	}
	return nil
}

// 保存角色错误事件处理
func (s *RoleServices) RoleSaveFailedEventHandle(event events.RoleSaveFailedEvent) error {
	s.logger.Sugar().Errorf("save event role failed %v", event.Err)
	return event.Err
}

// 删除角色事件处理
func (s *RoleServices) DeleteRoleEventHandle(event events.RoleDeletedEvent) error {
	if err := s.repo.DeleteRole(event.Id); err != nil {
		return errors.New("delete role failed " + err.Error())
	}
	return nil
}

// 删除角色错误事件处理
func (s *RoleServices) RoleDeleteFailedEventHandle(event events.RoleDeleteFailedEvent) error {
	s.logger.Sugar().Errorf("delete role failed %v", event.Err)
	return event.Err
}

// 获取角色列表
func (s *RoleServices) GetRoleList() ([]repository.Roles, error) {
	return s.repo.GetRoleList()
}