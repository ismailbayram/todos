package users

import (
	"crypto/sha256"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Username  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	IsActive  bool      `gorm:"default:true;not null"`
	IsAdmin   bool      `gorm:"default:false;not null"`
}

func CreateUser(db *gorm.DB, username string, password string, isAdmin bool) *User {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := h.Sum(nil)

	user := &User{
		Username: username,
		Password: string(hashedPassword),
		IsAdmin:  isAdmin,
		IsActive: true,
	}
	db.Create(user)
	return user
}
