package framework

import (
	"log"
	"net/http"
)

type Engine struct {
	Router *Router
}

func NewEngine() *Engine {
	return &Engine{
		Router: &Router{
			routingTable: Constructor(),
		},
	}
}

type Router struct {
	routingTable TreeNode
}

func (r *Router) Get(pathname string, handler func(w http.ResponseWriter, r *http.Request)) error {
	r.routingTable.Insert(pathname, handler)
	return nil
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		path := r.URL.Path
		handler := e.Router.routingTable.Search(path)
		if handler == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handler(w, r)
		return
	}
}

func (e *Engine) Run() {
	log.Fatalln(http.ListenAndServe(":8080", e))
}
