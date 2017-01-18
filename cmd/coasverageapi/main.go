package main

import (
	"github.com/udistrital/coasverage-api/internal/coasverageapi"
	"log"
	"net/http"
	"os"
)

func main() {
	var listen string
	var listen bool
	listen, listen_found = os.LookupEnv("LISTEN")
	if !listen_found {
		listen = "127.0.0.1:8000"
	}
	log.Fatal(http.ListenAndServe(listen_and_serve, coasverageapi.NewRouter()))
}
