package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-login/db"
	"go-login/mail"
	"net/http"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	mailer := mail.NewMailhogMailer()

	e := NewRouter(db, mailer)

	// error_handler.goの内容を登録してます。
	e.HTTPErrorHandler = customHTTPErrorHandler

	// validator.goの内容を登録してます。
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Logger.Fatal(e.Start(":8000"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func customHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)

	// エラーの内容をそのまま返すのは本当はNG
	if err := c.JSON(http.StatusInternalServerError, echo.Map{
		"message": err.Error(),
	}); err != nil {
		c.Logger().Error(err)
	}
}
