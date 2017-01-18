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
		Name:        "CoveragesGet",
		Method:      "GET",
		Pattern:     "/coverages",
		HandlerFunc: CoveragesGet,
	},
	Route{
		Name:        "CoverageGet",
		Method:      "GET",
		Pattern:     "/coverage/{id}",
		HandlerFunc: CoverageGet,
	},
	Route{
		Name:        "CoveragePost",
		Method:      "POST",
		Pattern:     "/coverage",
		HandlerFunc: CoveragePost,
	},
}
