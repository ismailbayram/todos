package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
	Username  string    `gorm:"not null;uniqueIndex:idx_name"`
	Password  string    `gorm:"not null"`
	IsActive  bool      `gorm:"default:true;not null"`
	IsAdmin   bool      `gorm:"default:false;not null"`
}
