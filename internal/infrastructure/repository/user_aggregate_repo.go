package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// UserAggregateRepository 用户聚合仓储
type UserAggregateRepository struct {
	db      *gorm.DB
	tx      *gorm.DB
	redis 	*redis.Client
	isAdmin bool
}

// NewUserAggregateRepository 创建用户聚合仓储实例
func NewUserAggregateRepository(db *gorm.DB, redisClient *redis.Client, isAdmin bool) user.IUserAggregateRepository {
	return &UserAggregateRepository{
		db:      db,
		tx:      nil,
		redis: redisClient,
		isAdmin: isAdmin,
	}
}

// dbInstance 获取数据库实例
func (r *UserAggregateRepository) dbInstance() *gorm.DB {
	db := r.db
	if r.tx != nil {
		db = r.tx
	}
	return db
}

// Begin 开启事务
func (r *UserAggregateRepository) Begin() *gorm.DB {
	r.tx = r.db.Begin()
	return r.tx
}

// IsAdmin 检查是否为管理员
func (r *UserAggregateRepository) IsAdmin() bool {
	return r.isAdmin
}

// ExistUserId 检查用户ID是否存在
func (r *UserAggregateRepository) ExistUserId(userId *string) bool {
	var count int64
	r.db.Model(&model.UserAggregate{}).Where("user_id = ?", userId).Count(&count)
	return count > 0
}

// ExistUser 检查用户名称是否存在
func (r *UserAggregateRepository) ExistUser(name *string) bool {
	var count int64
	r.db.Model(&model.UserAggregate{}).
		Where("username = ?", name).
		Count(&count)
	return count > 0
}

// ExistRole 检查角色是否存在
func (r *UserAggregateRepository) ExistRole(roleId *string) bool {
	var count int64

	r.db.Model(&model.Role{}).Where("id = ?", "da11f1bb-5727-4d79-bdcc-78e526e10362").Debug().Count(&count)
	return count > 0
}

// GetUserAggregate 获取用户聚合
func (r *UserAggregateRepository) GetUserAggregate(userId *string) (*user.UserAggregate, error) {
	var userModel model.UserAggregate
	if err := r.db.Where("user_id = ?", userId).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	return &user.UserAggregate{
		UserId:        &userModel.UserID,
		Username:      &userModel.Username,
		Status:        userModel.Status,
		PasswordHash:  &userModel.PasswordHash,
		Role:          userModel.Role,
		DeletedAt:     &userModel.DeletedAt.Time,
		Email:         userModel.Email,
		LastLoginAt:   userModel.LastLoginAt,
		LastLogoutAt:  userModel.LastLogoutAt,
		LoginFailedAt: userModel.LoginFailedAt,
		Attempts:      userModel.Attempts,
	}, nil
}

// SaveUserAggregate 保存用户聚合
func (r *UserAggregateRepository) SaveUserAggregate(aggregate *user.UserAggregate) error {
	db := r.dbInstance()

	var userModel model.UserAggregate
	// 更新非空字段
	userModel.UserID = *aggregate.UserId
	userModel.Username = derefOr(aggregate.Username, userModel.Username)
	userModel.PasswordHash = derefOr(aggregate.PasswordHash, userModel.PasswordHash)
	userModel.Status = aggregate.Status
	userModel.Role = aggregate.Role
	userModel.Email = aggregate.Email
	userModel.LastLoginAt = aggregate.LastLoginAt
	userModel.LastLogoutAt = aggregate.LastLogoutAt
	userModel.LoginFailedAt = aggregate.LoginFailedAt
	userModel.Attempts = aggregate.Attempts

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

// DeleteUserAggregate 删除用户聚合
func (r *UserAggregateRepository) DeleteUserAggregate(userId *string) error {
	db := r.dbInstance()
	return db.Delete(&model.UserAggregate{}, "user_id = ?", userId).Error
}

// derefOr 如果 ptr 不为 nil，返回 ptr 的值，否则返回 defaultValue
func derefOr[T any](ptr *T, defaultValue T) T {
	if ptr != nil {
		return *ptr
	}
	return defaultValue
}

// 验证邮箱验证码
func (r *UserAggregateRepository)	VerifyEmailCode(userId string, code string) error {
	rdb := r.redis
	val, err := rdb.Get(context.Background(), "email_valid_code:"+userId).Result()
	if err == redis.Nil {
		return errors.New("name does not exist")
	} else if err != nil {
		return errors.New("redis error" + err.Error())  
	} else {
		fmt.Println("name", val)
	}

	if val != code {
		return errors.New("code is not valid")
	}

	return nil
}

// 检查用户名是否已存在
func (r *UserAggregateRepository)	ExistsByUsername(username string) (bool, error) {
	var count int64
	if err := r.db.Model(&model.UserAggregate{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

// 检查邮箱是否已存在
func(r *UserAggregateRepository)	ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&model.UserAggregate{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	
	return count > 0 , nil
}

// 根据用户名获取用户聚合
func(r *UserAggregateRepository) GetUserAggregateByUsername(username string) (*user.UserAggregate, error) {
	var userAggregate model.UserAggregate
	if err := r.db.Where("username = ?", username).First(&userAggregate).Error; err != nil {
		return nil, err
	}
	
	return &user.UserAggregate{
		UserId: &userAggregate.UserID, 
		Username: &userAggregate.Username, 
		PasswordHash: &userAggregate.PasswordHash, 
		Role: userAggregate.Role,
		Status: userAggregate.Status,
		DeletedAt: &userAggregate.DeletedAt.Time,
		Email:         userAggregate.Email,
		LastLoginAt:   userAggregate.LastLoginAt,
		LastLogoutAt:  userAggregate.LastLogoutAt,
		LoginFailedAt: userAggregate.LoginFailedAt,
		Attempts:      userAggregate.Attempts,
	}, nil
}