package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/controllers"
	"go-gin/models"
	"go-gin/repositories"
	"go-gin/services"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	}

	itemRepository := repositories.NewItemRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()

	itemsRouter := r.Group("/items")
	itemsRouter.GET("/", itemController.FindAll)
	itemsRouter.GET("/:id", itemController.FindById)
	itemsRouter.POST("/", itemController.Create)
	itemsRouter.PUT("/:id", itemController.Update)
	itemsRouter.DELETE("/:id", itemController.Delete)

	r.Run()
}
