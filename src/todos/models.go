package todos

import (
	"github.com/ismailbayram/todos/src/users"
	"time"
)

type ToDo struct {
	ID        uint       `gorm:"primarykey"`
	CreatedAt time.Time  `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time  `gorm:"not null;autoUpdateTime"`
	Name      string     `gorm:"not null"`
	IsDone    bool       `gorm:"not null"`
	UserID    uint       `gorm:"not null"`
	User      users.User `gorm:"not null;foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
