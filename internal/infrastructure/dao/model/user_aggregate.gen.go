// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserAggregate = "user_aggregate"

// UserAggregate mapped from table <user_aggregate>
type UserAggregate struct {
	UserID        string         `gorm:"column:user_id;primaryKey" json:"user_id"`
	Username      string         `gorm:"column:username;not null" json:"username"`
	Email         *string        `gorm:"column:email" json:"email"`
	PasswordHash  string         `gorm:"column:password_hash;not null" json:"password_hash"`
	Role          *string        `gorm:"column:role" json:"role"`
	Status        *string        `gorm:"column:status" json:"status"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	LastLoginAt   *time.Time     `gorm:"column:last_login_at" json:"last_login_at"`
	LastLogoutAt  *time.Time     `gorm:"column:last_logout_at" json:"last_logout_at"`
	LoginFailedAt *time.Time     `gorm:"column:login_failed_at" json:"login_failed_at"`
	LoginAttempts *int32         `gorm:"column:login_attempts" json:"login_attempts"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName UserAggregate's table name
func (*UserAggregate) TableName() string {
	return TableNameUserAggregate
}
