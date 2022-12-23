package repositories

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/ismailbayram/todos/src/models"
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

func (r *UserRepository) CreateUser(username string, password string, isAdmin bool) (*models.User, error) {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := h.Sum(nil)

	user := &models.User{
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

func (r *UserRepository) ActivateUser(user *models.User) {
	user.IsActive = true
	r.db.Save(user)
}

func (r *UserRepository) DeactivateUser(user *models.User) error {
	user.IsActive = false
	return r.db.Save(user).Error
}
