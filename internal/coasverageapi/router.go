package coasverageapi

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

		var access_log string
		var access_log_found bool
		access_log, access_log_found = os.LookupEnv("ACCESS_LOG")
		if !access_log_found {
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
