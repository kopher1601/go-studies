package main

import (
	"go-tucker-web/web10/deco"
	"go-tucker-web/web10/myapp"
	"log"
	"net/http"
	"time"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Complete time", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Complete time", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	h := deco.NewDecoHandler(mux, logger)
	h = deco.NewDecoHandler(h, logger2)

	return h
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}
