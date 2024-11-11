package controllers

import (
	"gin-fleamarket/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ItermController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
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

func (i *ItemControllerImpl) FindByID(ctx *gin.Context) {
	itemID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item id"})
		return
	}
	item, err := i.service.FindById(uint(itemID))
	if err != nil {
		if err.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}
