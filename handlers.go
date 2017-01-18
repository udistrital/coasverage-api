package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CoveragesGet(w http.ResponseWriter, r *http.Request) {
	coverages, _ := ListCoverages()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err := json.NewEncoder(w).Encode(coverages); err != nil {
		panic(err)
	}
}

func CoverageGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coverage_id := vars["id"]
	fmt.Fprintf(w, "Coverage id: %v\n", coverage_id)
}

func CoveragePost(w http.ResponseWriter, r *http.Request) {
	var cov Coverage
	var err error
	json_decoder := json.NewDecoder(r.Body)
	if err = json_decoder.Decode(&cov); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else if err = UpdateCoverage(cov); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
