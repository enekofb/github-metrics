package metrics

import (
	"github.com/enekofb/metrics/pkg/config"
	"github.com/enekofb/metrics/pkg/issues"
	github2 "github.com/google/go-github/v48/github"
)

type QueryFunc func() (int, error)

func NewFromConfig(queryConfig config.QueryConfig) func() (int, error) {

	return func() (int, error) {

		owner := queryConfig.Owner
		if owner == "" {
			return -1, nil
		}

		repo := queryConfig.Repo
		if repo == "" {
			return -1, nil
		}

		defectLabel := queryConfig.BugLabel
		if defectLabel == "" {
			return -1, nil
		}

		teamLabel := queryConfig.TeamLabel
		if teamLabel == "" {
			return -1, nil
		}

		//oneDay := time.Hour * 24
		//lastMonth := time.Now().Add(-oneDay)

		options := &github2.IssueListByRepoOptions{
			Labels: []string{
				defectLabel,
				teamLabel,
			},
			//Since: lastMonth,
		}

		repoIssues, err := issues.GetRepoIssues(owner, repo, options)
		if err != nil {
			return -1, nil
		}
		return len(repoIssues), nil

	}
}
