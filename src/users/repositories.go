package users

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(username string, password string, isAdmin bool) (*User, error) {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := h.Sum(nil)

	user := &User{
		Username: username,
		Password: hex.EncodeToString(hashedPassword),
		IsAdmin:  isAdmin,
		IsActive: true,
	}
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) GetByID(id uint) (*User, error) {
	var user User
	result := r.db.Where("id = ?", id).First(&user)
	return &user, result.Error
}

func (r *UserRepository) GetByUsername(username string) (*User, error) {
	var user User
	result := r.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (r *UserRepository) CheckUserPassword(user *User, password string) bool {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := h.Sum(nil)
	return user.Password == hex.EncodeToString(hashedPassword)
}

func (r *UserRepository) Activate(user *User) error {
	user.IsActive = true
	return r.db.Save(user).Error
}

func (r *UserRepository) Deactivate(user *User) error {
	user.IsActive = false
	return r.db.Save(user).Error
}

func (r *UserRepository) MakeAdmin(user *User) error {
	user.IsAdmin = true
	return r.db.Save(user).Error
}

func (r *UserRepository) CreateToken(user *User) (string, error) {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", "token", viper.GetString("secretkey"))))
	token := hex.EncodeToString(h.Sum(nil))
	user.Token = token
	return token, r.db.Save(user).Error
}

func (r *UserRepository) GetByToken(token string) (*User, error) {
	var user User
	result := r.db.Where("token = ?", token).First(&user)
	return &user, result.Error
}
