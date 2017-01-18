package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCoverages(w http.ResponseWriter, r *http.Request) {
	var err error
	var coverages Coverages
	if coverages, err = ListCoverages(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err = json.NewEncoder(w).Encode(coverages); err != nil {
		panic(err)
	}
}

func GetCoverage(w http.ResponseWriter, r *http.Request) {
	var err error
	var coverage Coverage
	var vars map[string]string
	vars = mux.Vars(r)
	coverage_id := vars["id"]
	if coverage, err = ReadCoverage(coverage_id); err == nil {
		w.WriteHeader(http.StatusOK)
	} else if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, err.Error())
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err = json.NewEncoder(w).Encode(coverage); err != nil {
		panic(err)
	}
}

func PostCoverage(w http.ResponseWriter, r *http.Request) {
	var cov Coverage
	var err error
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1*1024*1024))
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &cov); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, err.Error())
	} else if err = UpdateCoverage(cov); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
