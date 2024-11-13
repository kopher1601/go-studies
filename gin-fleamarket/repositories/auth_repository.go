package repositories

import (
	"gin-fleamarket/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user models.User) error
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
