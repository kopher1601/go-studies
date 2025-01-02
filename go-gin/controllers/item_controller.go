package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin/services"
	"log"
	"net/http"
)

type ItemController interface {
	FindAll(ctx *gin.Context)
}

type itemController struct {
	service services.ItemService
}

func NewItemController(service services.ItemService) ItemController {
	return &itemController{service: service}
}

func (i *itemController) FindAll(ctx *gin.Context) {
	items, err := i.service.FindAll()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
