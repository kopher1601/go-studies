package controllers

import (
	"gin-fleamarket/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ItermController interface {
	FindAll(ctx *gin.Context)
}

type ItemControllerImpl struct {
	service services.ItemService
}

func NewItemController(service services.ItemService) ItermController {
	return &ItemControllerImpl{service: service}
}

func (i *ItemControllerImpl) FindAll(ctx *gin.Context) {
	items, err := i.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
