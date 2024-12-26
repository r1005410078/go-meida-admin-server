package command

import (
	"time"

	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
)

// 保存用户命令
type SaveUserCommand struct {
	ID                string            `json:"id"`                  // 用户唯一标识
	Username          string            `json:"username"`            // 用户名
	Email             string            `json:"email"`               // 邮箱
	Phone             string            `json:"phone"`               // 手机号
	FullName          string            `json:"full_name"`           // 全名
	AvatarURL         string            `json:"avatar_url"`          // 头像链接
	Gender            string            `json:"gender"`              // 性别
	Birthday          time.Time         `json:"birthday"`            // 出生日期
	Address           string            `json:"address"`             // 地址
	PasswordHash      string            `json:"password_hash"`      // 密码哈希
	Status            string            `json:"status"`              // 账户状态
	RoleId            string            `json:"role"`               // 用户角色
	Preferences       map[string]string `json:"preferences"`         // 用户偏好设置
	ReferredBy        string            `json:"referred_by"`         // 邀请者 ID
}
 
func (command *SaveUserCommand) ToEvent() *events.SaveUserEvent {
	return &events.SaveUserEvent{
		ID:                command.ID,
		Username:          command.Username,
		Email:             command.Email,
		Phone:             command.Phone,
		FullName:          command.FullName,
		AvatarURL:         command.AvatarURL,
		Gender:            command.Gender,
		Birthday:          command.Birthday,
		Address:           command.Address,
		PasswordHash:      command.PasswordHash,
		Status:            command.Status,
		RoleId:            command.RoleId,
		Preferences:       command.Preferences,
		ReferredBy:        command.ReferredBy,
	}
}