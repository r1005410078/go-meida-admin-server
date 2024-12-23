package main

import (
	"slices"

	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/db"
)


func main() {
	db, _ := db.GetDB()

	type Roles struct {
		model.Role
		Permissions []model.UserPermission `json:"permissions"`
	}

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

	rolesMap := make(map[string]Roles)
	permissionIds := []string{}

	for _, result := range results {
		if _, exists := rolesMap[result.ID]; !exists {
			rolesMap[result.ID] = Roles{result.Role, []model.UserPermission{}}
		}

		if !slices.Contains(permissionIds, result.PermissionId) {
			permissionIds = append(permissionIds, result.PermissionId)
		}

		rolesMap[result.ID] = Roles{result.Role, append(rolesMap[result.ID].Permissions, model.UserPermission{ID: result.PermissionId})}
	}

	UserPermissions := []model.UserPermission{}
	db.Model(&model.UserPermission{}).Where("id in ?", permissionIds).Find(&UserPermissions)

	permissionsMap := make(map[string]model.UserPermission)

	for _, permission := range UserPermissions {
		permissionsMap[permission.ID] = permission
	}

	for _, role := range rolesMap {
		for index, permission := range role.Permissions {
			role.Permissions[index] = permissionsMap[permission.ID]
		}
	}
}