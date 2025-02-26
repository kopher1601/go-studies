package main

import (
	"go-zero-to-one/controller"
	"go-zero-to-one/framework"
)

func main() {
	engine := framework.NewEngine()
	router := engine.Router

	router.Get("/lists", controller.ListController)
	router.Get("/lists/:list_id", controller.ListItemController)
	router.Get("/lists/:list_id/pictures/:picture_id", controller.ListItemPictureItemController)
	router.Get("/users", controller.UsersController)
	router.Get("/students", controller.StudentController)

	router.Get("/posts", controller.PostsPageController)
	router.Post("/posts", controller.PostsController)
	engine.Run()
}
