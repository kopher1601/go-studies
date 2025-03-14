package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-login/entity"
	"time"
)

type UserRepository interface {
	PreRegister(ctx context.Context, user *entity.User) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Delete(ctx context.Context, id entity.UserID) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

// ユーザーをstate=inactiveで保存する
func (u *userRepository) PreRegister(ctx context.Context, user *entity.User) error {
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()
	user.State = entity.UserInactive

	query := `INSERT INTO user(email, password, salt, activate_token, state, updated_at, created_at) 
VALUES (:email, :password, :salt, :activate_token, :state, :updated_at, :created_at)`
	result, err := u.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return fmt.Errorf("failed to Exec: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to GetLastInsertId: %w", err)
	}
	user.ID = entity.UserID(id)
	return nil
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, email, password, salt, state, activate_token, updated_at, created_at
FROM user WHERE email = ?`
	user := &entity.User{}
	if err := u.db.GetContext(ctx, user, query, email); err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return user, nil
}

func (u *userRepository) Delete(ctx context.Context, id entity.UserID) error {
	query := `DELETE FROM user WHERE id = ?`

	if _, err := u.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
