package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-gin/models"
	"go-gin/repositories"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService interface {
	Signup(email, password string) error
	Login(email, password string) (*string, error)
	GetUserFromToken(tokenString string) (*models.User, error)
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

func (a *authService) Login(email, password string) (*string, error) {
	foundUser, err := a.repository.FindUser(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	token, err := CreateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func CreateToken(useId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   useId,
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (a *authService) GetUserFromToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 서명 방법의 확인
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	var user *models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, jwt.ErrTokenExpired
		}

		user, err = a.repository.FindUser(claims["email"].(string))
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}
