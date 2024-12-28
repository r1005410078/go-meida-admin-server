package repository

import (
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"gorm.io/gorm"
)

// UserAggregateRepository 用户聚合仓储
type UserAggregateRepository struct {
	db      *gorm.DB
	tx      *gorm.DB
	isAdmin bool
}

// NewUserAggregateRepository 创建用户聚合仓储实例
func NewUserAggregateRepository(db *gorm.DB, isAdmin bool) user.IUserAggregateRepository {
	return &UserAggregateRepository{
		db:      db,
		isAdmin: isAdmin,
	}
}

func (r *UserAggregateRepository) dbInstance() *gorm.DB {
	db := r.db
	if r.tx != nil {
		db = r.tx
	}
	return db
}

func (r *UserAggregateRepository) Begin() *gorm.DB {
	r.tx = r.db.Begin()
	return r.tx
}

func (r *UserAggregateRepository) IsAdmin() bool {
	return r.isAdmin
}

func (r *UserAggregateRepository) ExistUserId(userId *string) bool {
	var count int64
	r.db.Model(&model.UserAggregate{}).Where("user_id = ?", userId).Count(&count)
	return count > 0
}

func (r *UserAggregateRepository) ExistUser(name *string) bool {
	var count int64
	r.db.Model(&model.UserAggregate{}).
		Where("username = ?", name).
		Count(&count)
	return count > 0
}

func (r *UserAggregateRepository) ExistRole(roleId *string) bool {
	var count int64

	r.db.Model(&model.Role{}).Where("id = ?", "da11f1bb-5727-4d79-bdcc-78e526e10362").Debug().Count(&count)
	return count > 0
}

func (r *UserAggregateRepository) GetUserAggregate(userId *string) (*user.UserAggregate, error) {
	var userModel model.UserAggregate
	if err := r.db.Where("user_id = ?", userId).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	return &user.UserAggregate{
		UserId:       &userModel.UserID,
		Username:     &userModel.Username,
		Status:       userModel.Status,
		PasswordHash: &userModel.PasswordHash,
		Role:         userModel.Role,
		DeletedAt:    &userModel.DeletedAt,
	}, nil
}

func (r *UserAggregateRepository) SaveUserAggregate(aggregate *user.UserAggregate) error {
	db := r.dbInstance()

	var userModel model.UserAggregate
	// 更新非空字段
	userModel.UserID = *aggregate.UserId
	userModel.Username = derefOr(aggregate.Username, userModel.Username)
	userModel.PasswordHash = derefOr(aggregate.PasswordHash, userModel.PasswordHash)
	userModel.Status = aggregate.Status
	userModel.Role = aggregate.Role

	// 使用 Updates 进行条件更新，如果记录不存在则创建
	result := db.Model(&model.UserAggregate{}).Where("user_id = ?", userModel.UserID).Updates(userModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return r.db.Create(&userModel).Error
	}

	return nil
}

func (r *UserAggregateRepository) DeleteUserAggregate(userId *string) error {
	db := r.dbInstance()
	return db.Delete(&model.UserAggregate{}, "user_id = ?", userId).Error
}

func derefOr[T any](ptr *T, defaultValue T) T {
	if ptr != nil {
		return *ptr
	}
	return defaultValue
}
