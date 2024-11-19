package main

import (
	"go-tucker-web/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", myapp.NewHttpHandler())
}
