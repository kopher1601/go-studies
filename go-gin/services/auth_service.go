package services

import (
	"go-gin/models"
	"go-gin/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Signup(email, password string) error
}

type authService struct {
	repository repositories.AuthRepository
}

func NewAuthService(repository repositories.AuthRepository) AuthService {
	return &authService{repository: repository}
}

func (a *authService) Signup(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return a.repository.CreateUser(user)
}