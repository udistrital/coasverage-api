package coasverageapi

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) (handler http.Handler) {
	handler_func := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
	}
	handler = http.HandlerFunc(handler_func)
	return
}
