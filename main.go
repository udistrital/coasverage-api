package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/coverages", CoverageIndex)
	router.HandleFunc("/coverage/{id}", Coverage)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func CoverageIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Coverage Index!")
}

func Coverage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coverage_id := vars["id"]
	fmt.Fprintf(w, "Coverage id: %v\n", coverage_id)
}
