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

func ListItemPictureItemController(ctx *framework.MyContext) {
	listID := ctx.GetParam(":list_id", "")
	pictureID := ctx.GetParam(":picture_id", "")

	output := struct {
		ListID    string `json:"list_id,omitempty"`
		PictureID string `json:"picture_id,omitempty"`
	}{
		ListID:    listID,
		PictureID: pictureID,
	}

	ctx.Json(&output)
}
