package repository

import (
	"context"
	"errors"

	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// UserRepository 实现用户仓储接口
type UserRepository struct {
	db *gorm.DB
	redisDb *redis.Client
	ctx context.Context
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository(db *gorm.DB, redisDb *redis.Client, ctx context.Context) *UserRepository {
	return &UserRepository{db: db, redisDb: redisDb, ctx: ctx}
}

// FindById 根据ID查找用户
func (r *UserRepository) FindById(id string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Save 保存用户信息
func (r *UserRepository) Save(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete 删除用户
func (r *UserRepository) Delete(user *model.User) error {
	return r.db.Delete(user).Error
}

// List 获取用户列表
func (r *UserRepository) List() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// AssoicatedRoles 关联角色
func (r *UserRepository) AssoicatedRoles(event *events.AssoicatedRolesEvent) error {
	if err := r.db.Model(&model.User{}).
		Where("id=?", event.UserId).
		Update("role", event.RoleId).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser 删除用户（根据事件）
func (r *UserRepository) DeleteUser(event *events.UserDeletedEvent) error {
	return r.db.Delete(&model.User{}, "id = ?", event.Id).Error
}

// SaveUser 保存用户（根据事件）
func (r *UserRepository) SaveUser(event *events.SaveUserEvent) error {
	if event.ID == nil {
		return errors.New("用户id不能为空")
	}

	user := &model.User{
		ID:        *event.ID,
		Email:     event.Email,
		Phone:     event.Phone,
		FullName:  event.FullName,
		AvatarURL: event.AvatarURL,
		Gender:    event.Gender,
		Address:   event.Address,
		Status:    event.Status,
		Role:      event.RoleId,
	}

	// 密码
	if event.PasswordHash != nil {
		user.PasswordHash = *event.PasswordHash
	}
	
	if event.Username != nil {
		user.Username = *event.Username
	}

	// 使用 Updates 进行条件更新，如果记录不存在则创建
	result := r.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	// 如果记录不存在（RowsAffected = 0），则创建新记录
	if result.RowsAffected == 0 {
		return r.db.Create(user).Error
	}

	return nil
}

// SaveUserStatus 保存用户状态
func (r *UserRepository) SaveUserStatus(event *events.UserStatusEvent) error {
	return r.db.Model(&model.User{}).
		Where("id = ?", event.Id).
		Update("status", event.Status).
		Error
}


// 根据邮箱获取用户
func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 保存Email验证码
func (r *UserRepository)  SaveEmailCode(email string, code string) error {
	return r.db.Model(&model.User{}).
		Where("email = ?", email).
		Update("email_code", code).
		Error
}

// 保存登陆token
func (r *UserRepository)  SaveLoginToken(userId string, token string) error {
	if err := r.redisDb.Set(r.ctx, "token:"+userId, token, 0).Err(); err != nil {
		return err
	}

	return nil
}
// 删除登陆token
func (r *UserRepository) DeleteLoginToken(userId *string) error {
	if err := r.redisDb.Del(r.ctx, "token:"+*userId).Err(); err != nil {
		return err
	}	

	return nil
}