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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err := json.NewEncoder(w).Encode(coverages); err != nil {
		panic(err)
	}
}

func CoverageShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coverage_id := vars["id"]
	fmt.Fprintf(w, "Coverage id: %v\n", coverage_id)
}
