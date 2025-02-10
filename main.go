package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Endpoint string
	Handler  func(http.ResponseWriter, *http.Request)
}

var Endpoints = []Route{
	{Endpoint: "/" , Handler: IndexHandler},
}


func main() {
	router := mux.NewRouter()
	fmt.Println("Starting server...")
	for _, route := range Endpoints {
		router.HandleFunc(route.Endpoint, route.Handler).Methods("GET")
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}



func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
