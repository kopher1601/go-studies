package repositories

import (
	"go-gin/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user models.User) error
	FindUser(email string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) CreateUser(user models.User) error {
	result := a.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *authRepository) FindUser(email string) (*models.User, error) {
	var user models.User
	result := a.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
