package main

import (
	"database/sql"
	"time"
)

type Coverage struct {
	Id               string        `json:"id"`
	UpdatedAt        time.Time     `json:"updated_at"`
	AppName          string        `json:"app_name"`
	RepoBranch       string        `json:"repo_name"`
	RepoCommit       string        `json:"repo_commit"`
	BuildEnvironment string        `json:"build_environment"`
	BuildCounter     int64         `json:"build_counter"`
	InternalBuildId  sql.NullInt64 `json:"internal_build_id"`
	CodeCoverage     float32       `json:"code_coverage"`
}

type Coverages []Coverage
