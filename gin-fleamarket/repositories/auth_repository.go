package repositories

import (
	"errors"
	"gin-fleamarket/models"
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
	result := a.db.First(&user, "email = ?", email)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}
