package router

import (
	"github.com/labstack/echo/v4"
	"go-echo/controller"
)

func NewRouter(uc controller.UserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.Signup)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	return e
}
