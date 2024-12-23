// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRolesPermission = "roles_permission"

// RolesPermission mapped from table <roles_permission>
type RolesPermission struct {
	RoleID       string    `gorm:"column:role_id;primaryKey" json:"role_id"`
	PermissionID string    `gorm:"column:permission_id;primaryKey" json:"permission_id"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName RolesPermission's table name
func (*RolesPermission) TableName() string {
	return TableNameRolesPermission
}
