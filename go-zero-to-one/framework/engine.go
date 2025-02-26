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
			routingTables: map[string]*TreeNode{
				"get":    Constructor(),
				"post":   Constructor(),
				"patch":  Constructor(),
				"put":    Constructor(),
				"delete": Constructor(),
			},
		},
	}
}

type Router struct {
	routingTables map[string]*TreeNode
}

func (r *Router) register(method string, pathname string, handler func(ctx *MyContext)) error {
	routingTable := r.routingTables[method]
	pathname = strings.TrimSuffix(pathname, "/")
	existedHandler := routingTable.Search(pathname)

	if existedHandler != nil {
		panic("already existed handler")
	}

	routingTable.Insert(pathname, handler)
	return nil
}

func (r *Router) Get(pathname string, handler func(ctx *MyContext)) error {
	return r.register("get", pathname, handler)
}

func (r *Router) Post(pathname string, handler func(ctx *MyContext)) error {
	return r.register("post", pathname, handler)
}

func (r *Router) Put(pathname string, handler func(ctx *MyContext)) error {
	return r.register("put", pathname, handler)
}

func (r *Router) Delete(pathname string, handler func(ctx *MyContext)) error {
	return r.register("delete", pathname, handler)
}

func (r *Router) Patch(pathname string, handler func(ctx *MyContext)) error {
	return r.register("patch", pathname, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewMyContext(w, r)
	routingTable := e.Router.routingTables[strings.ToLower(r.Method)]

	path := r.URL.Path
	path = strings.TrimSuffix(path, "/")
	targetNode := routingTable.Search(path)

	if targetNode == nil || targetNode.handler == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	paramDicts := targetNode.ParseParams(r.URL.Path)
	ctx.SetParams(paramDicts)
	targetNode.handler(ctx)
	return
}

func (e *Engine) Run() {
	log.Fatalln(http.ListenAndServe(":8080", e))
}
