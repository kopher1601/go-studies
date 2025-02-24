package framework

import (
	"errors"
	"log"
	"net/http"
)

type Engine struct {
	Router *Router
}

func NewEngine() *Engine {
	return &Engine{
		Router: &Router{},
	}
}

type Router struct {
	routingTable map[string]func(http.ResponseWriter, *http.Request)
}

func (r *Router) Get(pathname string, handler func(w http.ResponseWriter, r *http.Request)) error {
	if r.routingTable == nil {
		r.routingTable = make(map[string]func(http.ResponseWriter, *http.Request))
	}

	if r.routingTable[pathname] != nil {
		return errors.New("existed")
	}
	r.routingTable[pathname] = handler
	return nil
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handler := e.Router.routingTable[r.URL.Path]
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
