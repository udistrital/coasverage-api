package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Coverage struct {
	Id               string    `json:"id"`
	UpdatedAt        time.Time `json:"updated_at"`
	AppName          string    `json:"app_name"`
	RepoBranch       string    `json:"repo_name"`
	RepoCommit       string    `json:"repo_commit"`
	BuildEnvironment string    `json:"build_environment"`
	BuildCounter     int64     `json:"build_counter"`
	InternalBuildId  int64     `json:"internal_build_id"`
	CodeCoverage     float32   `json:"code_coverate"`
}

type Coverages []Coverage

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/coverages", CoverageIndex)
	router.HandleFunc("/coverage/{id}", CoverageShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}

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
