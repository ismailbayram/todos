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

func (r *UserRepository) Create(username string, password string, isAdmin bool) (*models.User, error) {
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

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.Where("id = ?", id).First(&user)
	return &user, result.Error
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (r *UserRepository) Activate(user *models.User) error {
	user.IsActive = true
	return r.db.Save(user).Error
}

func (r *UserRepository) Deactivate(user *models.User) error {
	user.IsActive = false
	return r.db.Save(user).Error
}

func (r *UserRepository) MakeAdmin(user *models.User) error {
	user.IsAdmin = true
	return r.db.Save(user).Error
}
