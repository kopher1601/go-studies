package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	//router := http.NewServeMux()
	router := mux.NewRouter()

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post request received")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
