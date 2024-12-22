package repository

import (
	"github.com/r1005410078/meida-admin-server/internal/domain/role/events"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"gorm.io/gorm"
)

// IRoleRepository 角色仓储接口
type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db,
	}
}

// 保存角色
func (r *RoleRepository) SaveRole(inputRole events.RoleSavedEvent) error {
	tx := r.db.Begin()
	role := model.Role {
		Name: inputRole.Name,
		Description: inputRole.Description,
	}

	// 角色是否存在
	var count int64 = 0
	if err := tx.Model(&model.Role{}).Where("id = ?", inputRole.Id).Count(&count).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 保存角色
	if count == 0 {
		role.ID = *inputRole.Id
		if err := tx.Create(&role).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := tx.Model(&model.Role{}).Where("id=?", inputRole.Id).Updates(&role).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 保存角色权限
	for _, PermissionId := range inputRole.PermissionIds {
		if err := tx.Create(&model.RolesPermission{ RoleID: *inputRole.Id, PermissionID: PermissionId}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	
	return nil
}

// 删除角色
func (r *RoleRepository) DeleteRole(id string) error {
	r.db.Delete(&model.Role{}, id)
	return nil
}
