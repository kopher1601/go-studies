package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StudentResponse struct {
	Name string `json:"name"`
}

func StudentController(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	name := queries.Get("name")

	studentResponse := &StudentResponse{
		Name: name,
	}
	responseData, err := json.Marshal(studentResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
	return
}

func ListController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "list")
}

func UsersController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "user")
}

func ListItemController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "list item")
}
