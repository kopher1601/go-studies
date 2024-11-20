package main

import (
	"go-tucker-web/web5/app"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", app.NewHandler())
}
