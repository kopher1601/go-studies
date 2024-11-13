package controllers

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController interface {
	Signup(ctx *gin.Context)
}

func NewAuthController(service services.AuthService) AuthController {
	return &authController{service: service}
}

type authController struct {
	service services.AuthService
}

func (a *authController) Signup(ctx *gin.Context) {
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := a.service.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}
