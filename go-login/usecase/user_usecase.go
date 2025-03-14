package usecase

import (
	"context"
	"database/sql"
	"errors"
	"go-login/entity"
	"go-login/mail"
	"go-login/repository"
	"math/rand"
)

type Usecase interface {
	PreRegister(ctx context.Context, email, password string) (*entity.User, error)
}

type userUsecase struct {
	ur     repository.UserRepository
	mailer mail.Mailer
}

func NewUserUsecase(ur repository.UserRepository, mailer mail.Mailer) Usecase {
	return &userUsecase{
		ur:     ur,
		mailer: mailer,
	}
}

func (u *userUsecase) PreRegister(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := u.ur.GetByEmail(ctx, email)

	// ユーザーが存在しない場合、sql.ErrNoRowsを受け取るはずなので、存在しない場合はそのまま仮登録処理を行う
	if errors.Is(err, sql.ErrNoRows) {
		return u.preRegister(ctx, email, password)
	} else if err != nil {
		return nil, err
	}

	// ユーザーがすでにアクティブの場合はエラーを返す
	if user.IsActive() {
		return nil, errors.New("user already active")
	}

	// ユーザーがアクティブではない場合、ユーザーを削除して、再度仮登録処理を行う
	if err := u.ur.Delete(ctx, user.ID); err != nil {
		return nil, err
	}
	return u.PreRegister(ctx, email, password)
}

func (u *userUsecase) preRegister(ctx context.Context, email string, password string) (*entity.User, error) {
	salt := createRandomString(30)
	activeToken := createRandomString(8)

	user := &entity.User{}

	//
	hashed, err := user.CreateHashedPassword(password, salt)
	if err != nil {
		return nil, err
	}

	user.Email = email
	user.Salt = salt
	user.Password = hashed
	user.ActivateToken = activeToken
	user.State = entity.UserInactive

	// DBへの仮登録処理を行う
	if err := u.ur.PreRegister(ctx, u); err != nil {
		return nil, err
	}
	// email宛に、本人確認用のトークンを送信する
	if err := u.mailer.SendWithActivateToken(email, user.ActivateToken); err != nil {
		return nil, err
	}
	return user, err
}

// lengthの長さのランダムな文字列(a-zA-Z0-9)を作成する
func createRandomString(length uint) string {
	var letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
