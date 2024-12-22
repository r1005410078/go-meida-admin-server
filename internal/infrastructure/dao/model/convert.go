package model

import (
	"encoding/json"

	"github.com/r1005410078/meida-admin-server/internal/domain/role"
)

func (m RoleAggregate) ToRoleAggregate() (*role.RoleAggregate, error) {
	PermissionIds := []string{}
	err := json.Unmarshal([]byte(m.PermissionIds), &PermissionIds)

	if err != nil {
		return nil, err
	}

	return &role.RoleAggregate{
		Id:        &m.ID,
		Name:      m.Name,
		DeletedAt: nil,
		UpdatedAt: &m.UpdatedAt,
		PermissionIds: PermissionIds,
	}, nil
}