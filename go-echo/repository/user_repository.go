package repository

import (
	"go-echo/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := u.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) CreateUser(user *model.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
