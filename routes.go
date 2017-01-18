package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "CoverageIndex",
		Method:      "GET",
		Pattern:     "/coverages",
		HandlerFunc: CoverageIndex,
	},
	Route{
		Name:        "Coverage",
		Method:      "GET",
		Pattern:     "/coverage/{id}",
		HandlerFunc: CoverageShow,
	},
}

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return
}
