package coasverageapi

import (
	"database/sql"
	"time"
)

func ListCoverages() (covs Coverages, err error) {
	var rows *sql.Rows
	if rows, err = ListCoveragesStmt.Query(); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var updated_at time.Time
		var app_name string
		var repo_branch string
		var repo_commit string
		var build_environment string
		var build_counter int64
		var internal_build_id sql.NullInt64
		var code_coverage float32
		if err = rows.Scan(&id, &updated_at, &app_name, &repo_branch, &repo_commit, &build_environment, &build_counter, &internal_build_id, &code_coverage); err == nil {
			covs = append(covs, Coverage{
				Id:               id,
				UpdatedAt:        updated_at,
				AppName:          app_name,
				RepoBranch:       repo_branch,
				RepoCommit:       repo_commit,
				BuildEnvironment: build_environment,
				BuildCounter:     build_counter,
				InternalBuildId:  internal_build_id,
				CodeCoverage:     code_coverage,
			})
		} else {
			return
		}
	}
	return
}
