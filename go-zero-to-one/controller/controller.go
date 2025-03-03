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

type PostPageForm struct {
	Name string
}

func PostsPageController(ctx *framework.MyContext) {
	authUser := ctx.Get("AuthUser", "defaultValue")
	postPageForm := &PostPageForm{
		Name: authUser.(string),
	}
	ctx.RenderHtml("./htmls/posts_page.html", postPageForm)
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

func JsonPTestController(ctx *framework.MyContext) {
	queryKey := ctx.QueryKey("callback", "cb")
	ctx.JsonP(queryKey, "")
}
