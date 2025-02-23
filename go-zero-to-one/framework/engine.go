package framework

import (
	"go-zero-to-one/controller"
	"log"
	"net/http"
)

type Engine struct {
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path == "/students" {
			controller.StudentController(w, r)
			return
		}

		if r.URL.Path == "/lists" {
			controller.ListController(w, r)
			return
		}

		if r.URL.Path == "/users" {
			controller.UsersController(w, r)
			return
		}

	}
}

func (e *Engine) Run() {
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
