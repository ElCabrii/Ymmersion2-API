package main

import "net/http"

type Route struct {
	Endpoint string
	Handler  func(http.ResponseWriter, *http.Request)
}

var Endpoints = []Route{
	{Endpoint: "/" , Handler: IndexHandler},
}
