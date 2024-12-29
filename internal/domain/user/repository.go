package user

import (
	"gorm.io/gorm"
)

type IUserAggregateRepository interface {
	// 开启事务
	Begin() *gorm.DB

	// 是否是管理员
	IsAdmin() bool

	// 用户ID是否存在
	ExistUserId(userId *string) bool

	// 用户名称是否存在
	ExistUser(name *string) bool

	// 是否存在角色
	ExistRole(roleId *string) bool

	// 获取聚合
	GetUserAggregate(userId *string) (*UserAggregate, error)

	// 保存聚合
	SaveUserAggregate(aggregate *UserAggregate) error

	// 删除聚合
	DeleteUserAggregate(userId *string) error

	// 验证邮箱验证码
	VerifyEmailCode(userId string, code string) error

	// 检查用户名是否已存在
	ExistsByUsername(username string) (bool, error)

	// 检查邮箱是否已存在
	ExistsByEmail(email string) (bool, error)

	// 获取用户信息
	GetUserAggregateByUsername(username string) (*UserAggregate, error)

}

