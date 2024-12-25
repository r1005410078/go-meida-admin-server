package permissions

import (
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
)

type PermissionsService struct {
	repo PermissionsRepositoryer
}

func NewPermissionsService(repo PermissionsRepositoryer) *PermissionsService {
	return &PermissionsService{
		repo,
	}
}

// 保存权限
func (s *PermissionsService) Save(permission *Permission) error {
	return s.repo.Save(permission);
}

// 删除权限
func (s *PermissionsService) Delete(permission *Permission) error {
	return s.repo.Delete(permission);
}

// 权限列表
func (s *PermissionsService) List()  ([]*model.UserPermission, error) {
	return s.repo.List();
}

