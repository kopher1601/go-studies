package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"go-login/handler"
	"go-login/mail"
	"go-login/repository"
	"go-login/usecase"
)

func NewRouter(db *sqlx.DB, mailer mail.Mailer) *echo.Echo {
	e := echo.New()

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur, mailer)
	uh := handler.NewUserHandler(uu)

	a := e.Group("/api/auth")
	a.POST("/register/initial", uh.PreRegister)

	return e
}
