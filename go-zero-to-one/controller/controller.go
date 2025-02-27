package controller

import (
	"fmt"
	"go-zero-to-one/framework"
	"io/fs"
	"net/http"
	"os"
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

func PostsController(ctx *framework.MyContext) {
	name := ctx.FormKey("name", "defaultName")
	age := ctx.FormKey("age", "20")
	fileInfo, err := ctx.FormFile("file")

	if err != nil {
		ctx.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(fmt.Sprintf("%s_%s_%s", name, age, fileInfo.Filename), fileInfo.Data, fs.ModePerm)
	if err != nil {
		ctx.WriteHeader(http.StatusInternalServerError)
	}
	ctx.WriteString("success")
}

func PostsPageController(ctx *framework.MyContext) {
	ctx.WriteString(`<!DOCTYPE html>
	<html>
		<head>
			<title>form</title>
		</head>
		<body>
			<div>
				<form action="/posts" method="post" enctype="multipart/form-data">
					<div><label>name</label>: <input name="name"/></div>
					<div><label>age</label>: 
					<select name="age">
						<option value="1">1</option>
						<option value="2">2</option>
					</select></div>
					<button type="submit">submit</button>
					<input name="file" type="file"/>
				</form>
			</div>
		</body>
	</html>`)
}

type UserPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UserController(ctx *framework.MyContext) {
	u := &UserPost{}
	if err := ctx.BindJson(u); err != nil {
		ctx.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Json(u)
}
