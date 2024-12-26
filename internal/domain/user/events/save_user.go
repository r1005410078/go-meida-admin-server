package events

import "time"

type SaveUserEvent struct {
	ID                string                   // 用户唯一标识
	Username          string          	       // 用户名
	Email             string                   // 邮箱
	Phone             string                   // 手机号
	FullName          string                   // 全名
	AvatarURL         string                   // 头像链接
	Gender            string                   // 性别
	Birthday          time.Time                // 出生日期
	Address           string                   // 地址
	PasswordHash      string                   // 密码哈希
	RoleId            string                   // 用户角色
	Status            string                   // 账户状态
	LastLoginAt       time.Time                // 最后登录时间
	LoginAttempts     int                      // 登录失败次数
	Preferences       map[string]string        // 用户偏好设置
	Tags              []string           		   // 用户标签
	ReferredBy        string                   // 邀请者 ID
}

type SaveUserFailedEvent struct {
	SaveUserEvent
	Err error
}

func NewSaveUserFailedEvent(SaveUserEvent *SaveUserEvent, Err error) SaveUserFailedEvent {
	return SaveUserFailedEvent{
		*SaveUserEvent,
		Err,
	}
}