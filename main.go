package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Uuid [16]byte

type Coverage struct {
	Id               Uuid
	UpdatedAt        time.Time
	AppName          string
	RepoBranch       string
	RepoCommit       string
	BuildEnvironment string
	BuildCounter     int64
	InternalBuildId  int64
	CodeCoverage     int
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
