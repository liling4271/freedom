package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         string         `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string         `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password   string         `gorm:"size:255;not null" json:"-"`
	UserStatus int8           `gorm:"size:1;default:1" json:"user_status"`
	LastLogin  *time.Time     `json:"last_login"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
