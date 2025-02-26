package controller

import (
	"go-zero-to-one/framework"
)

type StudentResponse struct {
	Name string `json:"name"`
}

func StudentController(ctx *framework.MyContext) {
	name := ctx.QueryKey("name", "")

	studentResponse := &StudentResponse{
		Name: name,
	}
	ctx.Json(studentResponse)
	return
}

func ListController(ctx *framework.MyContext) {
	ctx.WriteString("list")
}

func UsersController(ctx *framework.MyContext) {
	ctx.WriteString("user")
}

func ListItemController(ctx *framework.MyContext) {
	ctx.WriteString("list item")
}
