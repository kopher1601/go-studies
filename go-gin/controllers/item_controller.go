package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin/dto"
	"go-gin/services"
	"log"
	"net/http"
	"strconv"
)

type ItemController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
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

func (i *itemController) FindById(ctx *gin.Context) {
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	item, err := i.service.FindById(uint(itemId))
	if err != nil {
		if err.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

func (i *itemController) Create(ctx *gin.Context) {
	var input dto.CreateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem, err := i.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}
