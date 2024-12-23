package repository

import (
	"slices"

	"github.com/r1005410078/meida-admin-server/internal/app/repository"
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

func (r *RoleRepository) GetRoleList() ([]repository.Roles, error) {
	db := r.db

	var results []struct {
		model.Role
		PermissionId string
	}

  db.Table("roles").
		Raw(`
		SELECT 
			id, name, description, permission_id
			FROM roles r
			LEFT JOIN roles_permission rp ON r.id = rp.role_id
	`).Scan(&results)

	rolesMap := make(map[string]repository.Roles)
	permissionIds := []string{}

	for _, result := range results {
		if _, exists := rolesMap[result.ID]; !exists {
			rolesMap[result.ID] = repository.Roles{Role: result.Role, Permissions: []model.UserPermission{}}
		}

		if !slices.Contains(permissionIds, result.PermissionId) {
			permissionIds = append(permissionIds, result.PermissionId)
		}

		rolesMap[result.ID] = repository.Roles{Role: result.Role, Permissions: append(rolesMap[result.ID].Permissions, model.UserPermission{ID: result.PermissionId})}
	}

	UserPermissions := []model.UserPermission{}
	db.Model(&model.UserPermission{}).Where("id in ?", permissionIds).Find(&UserPermissions)

	permissionsMap := make(map[string]model.UserPermission)

	for _, permission := range UserPermissions {
		permissionsMap[permission.ID] = permission
	}

	roles := []repository.Roles{}
	for _, role := range rolesMap {
		for index, permission := range role.Permissions {
			role.Permissions[index] = permissionsMap[permission.ID]
		}
		roles = append(roles, role)
	}

	return roles, nil
}