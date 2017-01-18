package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CoverageIndex(w http.ResponseWriter, r *http.Request) {
	coverages := Coverages{
		Coverage{AppName: "app1"},
		Coverage{AppName: "app2"},
	}
	json.NewEncoder(w).Encode(coverages)
}

func CoverageShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coverage_id := vars["id"]
	fmt.Fprintf(w, "Coverage id: %v\n", coverage_id)
}
