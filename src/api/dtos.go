package api

import (
	"github.com/ismailbayram/todos/src/users"
	"time"
)

type LoginDTO struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type UserDTO struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
	IsAdmin   bool      `json:"is_admin"`
}

func ToUserDTO(user *users.User) UserDTO {
	return UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		IsActive:  user.IsActive,
		IsAdmin:   user.IsAdmin,
	}
}
