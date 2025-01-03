package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/controllers"
	"go-gin/infra"
	"go-gin/repositories"
	"go-gin/services"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	r := gin.Default()

	// items
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	itemsRouter := r.Group("/items")
	itemsRouter.GET("/", itemController.FindAll)
	itemsRouter.GET("/:id", itemController.FindById)
	itemsRouter.POST("/", itemController.Create)
	itemsRouter.PUT("/:id", itemController.Update)
	itemsRouter.DELETE("/:id", itemController.Delete)

	// auth
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.Signup)

	r.Run()
}
