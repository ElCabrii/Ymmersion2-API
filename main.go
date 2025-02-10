package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("Starting server...")
	for _, route := range Endpoints {
		router.HandleFunc(route.Endpoint, route.Handler).Methods("GET")
	}
	log.Fatal(http.ListenAndServe(":8000", router))
}
