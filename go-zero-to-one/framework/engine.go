package framework

import (
	"log"
	"net/http"
	"strings"
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

func (r *Router) Get(pathname string, handler func(ctx *MyContext)) error {
	existedHandler := r.routingTable.Search(pathname)

	if existedHandler != nil {
		panic("already existed handler")
	}

	r.routingTable.Insert(pathname, handler)
	return nil
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewMyContext(w, r)
	if r.Method == http.MethodGet {
		path := r.URL.Path
		path = strings.TrimSuffix(path, "/")
		targetNode := e.Router.routingTable.Search(path)

		if targetNode == nil || targetNode.handler == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		paramDicts := targetNode.ParseParams(r.URL.Path)
		ctx.SetParams(paramDicts)
		targetNode.handler(ctx)
		return
	}
}

func (e *Engine) Run() {
	log.Fatalln(http.ListenAndServe(":8080", e))
}
