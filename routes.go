package main

import (
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
