package main

import (
	"database/sql"
	"errors"
	"time"
)

var CreateCoverage = UpdateCoverage

func ReadCoverage(id string) (cov Coverage, err error) {
	var updated_at time.Time
	var app_name string
	var repo_branch string
	var repo_commit string
	var build_environment string
	var build_counter int64
	var internal_build_id sql.NullInt64
	var code_coverage float32
	if err = ReadCoverageStmt.QueryRow(id).Scan(&id, &updated_at, &app_name, &repo_branch, &repo_commit, &build_environment, &build_counter, &internal_build_id, &code_coverage); err == nil {
		cov.Id = id
		cov.UpdatedAt = updated_at
		cov.AppName = app_name
		cov.RepoBranch = repo_branch
		cov.RepoCommit = repo_commit
		cov.BuildEnvironment = build_environment
		cov.BuildCounter = build_counter
		cov.InternalBuildId = internal_build_id
		cov.CodeCoverage = code_coverage
	}
	return
}

func UpdateCoverage(cov Coverage) (err error) {
	if cov.InternalBuildId.Valid {
		internal_build_id := cov.InternalBuildId.Int64
		_, err = UpdateCoverageStmt.Exec(cov.AppName, cov.RepoBranch, cov.RepoCommit, cov.BuildEnvironment, cov.BuildCounter, internal_build_id, cov.CodeCoverage)
	} else {
		_, err = UpdateCoverageStmt.Exec(cov.AppName, cov.RepoBranch, cov.RepoCommit, cov.BuildEnvironment, cov.BuildCounter, nil, cov.CodeCoverage)
	}
	return
}

func DeleteCoverage(id string) (err error) {
	err = errors.New("not implemented")
	return
}
