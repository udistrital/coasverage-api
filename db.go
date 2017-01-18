package main

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var DB *sql.DB
var ReadCoverageStmt *sql.Stmt
var UpdateCoverageStmt *sql.Stmt

const read_coverage_query = `
SELECT id,
       updated_at,
       app_name,
       repo_branch,
       repo_commit,
       build_environment,
       build_counter,
       internal_build_id,
       code_coverage
FROM coasverage_model.coverages
WHERE id = $1`

const update_coverage_query = `
INSERT INTO coasverage_model.coverages ( updated_at, app_name, repo_branch, repo_commit, build_environment, build_counter, code_coverage )
VALUES ( NOW(),
         $1,
         $2,
         $3,
         $4,
         $5,
         $6 ) ON CONFLICT ( app_name, repo_branch, build_environment ) DO
UPDATE
SET updated_at = EXCLUDED.updated_at,
    repo_commit = EXCLUDED.repo_commit,
    build_counter = EXCLUDED.build_counter,
    internal_build_id = EXCLUDED.internal_build_id,
    code_coverage = EXCLUDED.code_coverage`

func init() {
	var postgres_url string
	var postgres_url_found bool
	var err error
	postgres_url, postgres_url_found = os.LookupEnv("POSTGRES_URL")
	if !postgres_url_found {
		panic(errors.New("no $POSTGRES_URL"))
	}
	if DB, err = sql.Open("postgres", postgres_url); err != nil {
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	log.Print("DB open")
	if ReadCoverageStmt, err = DB.Prepare(read_coverage_query); err != nil {
		panic(err)
	}
	if UpdateCoverageStmt, err = DB.Prepare(update_coverage_query); err != nil {
		panic(err)
	}
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		end()
	}()
}

func end() {
	var err error
	log.Print("DB closing")
	if err = ReadCoverageStmt.Close(); err != nil {
		log.Print(err.Error())
	}
	if err = UpdateCoverageStmt.Close(); err != nil {
		log.Print(err.Error())
	}
	if err = DB.Close(); err != nil {
		log.Print(err.Error())
	}
	os.Exit(0)
}
