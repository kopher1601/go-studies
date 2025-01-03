package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-gin/controllers"
	"go-gin/infra"
	"go-gin/middlewares"
	"go-gin/repositories"
	"go-gin/services"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	r := gin.Default()
	r.Use(cors.Default())

	// items
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	// auth
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	itemsRouter := r.Group("/items")
	itemsRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	itemsRouter.GET("/", itemController.FindAll)
	itemsRouterWithAuth.GET("/:id", itemController.FindById)
	itemsRouterWithAuth.POST("/", itemController.Create)
	itemsRouterWithAuth.PUT("/:id", itemController.Update)
	itemsRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	r.Run()
}
