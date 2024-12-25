package repository

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/r1005410078/meida-admin-server/internal/domain/role"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"gorm.io/gorm"
)

type RoleAggregateRepository  struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewRoleAggregateRepository(db *gorm.DB) *RoleAggregateRepository {
	return &RoleAggregateRepository{
		db: db,
		tx: nil,
	}
}

// 开启事务
func (r *RoleAggregateRepository) Begin() *gorm.DB {
	r.tx = r.db.Begin()
	return r.tx
}

func (r *RoleAggregateRepository) dbInstance() *gorm.DB {
	db := r.db
	if r.tx != nil {
		db = r.tx
	}
	return db
}

// 实现 IRoleAggregateRepository 接口
func (r *RoleAggregateRepository) IsAdmin() bool {
	return true
}

func (r *RoleAggregateRepository) IsValidPermissionID(id string) bool {
	return true
}

// 保存角色聚合
func (r *RoleAggregateRepository) SaveRoleAggregate(aggregate *role.RoleAggregate) error {
	db := r.dbInstance()

	if aggregate == nil {
		return errors.New("role aggregate is nil")
	}

  PermissionIds, err :=	json.Marshal(aggregate.PermissionIds)
	if err != nil {
		return err
	}

	roleModel := model.RoleAggregate {
		RoleID: *aggregate.RoleId,
		RoleName: aggregate.RoleName,
		PermissionIds: string(PermissionIds),
	}

	// 检查角色是否存在
	var count int64
	if err := db.Model(&model.RoleAggregate{}).Where("role_id = ?", roleModel.RoleID).Count(&count).Error; err != nil {
		return err
	}

	// 创建聚合
	if count == 0 {
		if result := db.Create(&roleModel); result.Error != nil {
			return result.Error
		}
		return nil
	}

	// 更新聚合
	if err := db.Model(&model.RoleAggregate{}).Where("role_id = ?", aggregate.RoleId).Updates(&roleModel).Error; err != nil {
		return err
	}
	
	return nil
}

func (r *RoleAggregateRepository) DeleteRoleAggregate(id string) error {
	r.db.Where("role_id=?", id).Delete(&model.RoleAggregate{})
	return nil
}

func (r *RoleAggregateRepository) GetRoleAggregate(id string) (*role.RoleAggregate, error) {
	var roleModel model.RoleAggregate

	if err := r.db.Where("role_id=?", id).First(&roleModel).Error; err != nil {
		return nil, err
	}

	res, err := roleModel.ToRoleAggregate()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *RoleAggregateRepository) ExistsPermissionIds(ids []string) bool {
	// 检查权限是否存在
	var exists bool
	if res := r.db.Model(&model.UserPermission{}).Select("count(*) > 0").Where("id in ?", ids).Find(&exists); res.Error != nil {
		log.Printf("ExistsPermissionIds error: %v", res.Error)
		return false
	}

	return exists
}

func (r *RoleAggregateRepository) IsRoleNameExist(name string) bool {
	if name == "" {
		return false
	}

	var exists bool
	r.db.Model(&model.RoleAggregate{}).Select("count(*) > 0").Where("name = ?", name).Find(&exists)

	return exists
}


