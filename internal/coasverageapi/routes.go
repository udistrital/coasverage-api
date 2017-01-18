package coasverageapi

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
		Name:        "GetCoverages",
		Method:      "GET",
		Pattern:     "/coverages",
		HandlerFunc: GetCoverages,
	},
	Route{
		Name:        "GetCoverage",
		Method:      "GET",
		Pattern:     "/coverage/{id}",
		HandlerFunc: GetCoverage,
	},
	Route{
		Name:        "PostCoverage",
		Method:      "POST",
		Pattern:     "/coverage",
		HandlerFunc: PostCoverage,
	},
}
