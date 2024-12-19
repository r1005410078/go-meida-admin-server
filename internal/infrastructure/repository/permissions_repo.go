package repository

import (
	"context"

	"github.com/r1005410078/meida-admin-server/internal/domain/permissions"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/query"
	"gorm.io/gorm"
)


type PermissionsRepository struct {
	db *gorm.DB
}

func NewPermissionsRepository(db *gorm.DB) *PermissionsRepository {
	return &PermissionsRepository{
		db,
	}
}

func (r *PermissionsRepository) Save(permission *permissions.Permission) error {
	newPermission := model.UserPermission { 
		Name: permission.Name,
		Description: permission.Description,
		Action: permission.Action,
	}

	p := query.Use(r.db).UserPermission.WithContext(context.Background());

	if permission.ID != nil {
		newPermission.ID = *permission.ID
		p.Updates(&newPermission)
	} else {
		p.Create(&newPermission)
	}

	return nil
}

func (r *PermissionsRepository) Delete(permission *permissions.Permission) error {
	return nil
}

func (r *PermissionsRepository) List() ([]*model.UserPermission, error) {
	list, err := query.Use(r.db).UserPermission.WithContext(context.Background()).Find()
	return list, err
}