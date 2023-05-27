package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type data struct {
	username int
	password int
	name     string
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my name is mahdi :)")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getRequest).Methods("GET")

	http.Handle("/", r)
	fmt.Println("server started and listening on localhost:9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}
