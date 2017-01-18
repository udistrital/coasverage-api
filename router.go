package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc

		access_log := os.Getenv("ACCESS_LOG")
		if access_log == "" {
			access_log = "yes"
		}
		if access_log == "yes" {
			handler = Logger(handler, route.Name)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return
}
