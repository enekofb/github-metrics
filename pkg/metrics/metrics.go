package metrics

import (
	"github.com/enekofb/metrics/pkg/config"
	"github.com/enekofb/metrics/pkg/issues"
	github2 "github.com/google/go-github/v48/github"
	"time"
)

func GetLastMonthDefectMetricsByTeam(config config.MetricsConfig) (int, error) {

	owner := config.GithubConfig.Owner
	if owner == "" {
		return -1, nil
	}

	repo := config.GithubConfig.Repo
	if repo == "" {
		return -1, nil
	}

	defectLabel := config.GithubConfig.BugLabel
	if defectLabel == "" {
		return -1, nil
	}

	teamLabel := config.GithubConfig.TeamLabel
	if teamLabel == "" {
		return -1, nil
	}

	oneMonth := time.Hour * 24 * 30
	lastMonth := time.Now().Add(-oneMonth)

	options := &github2.IssueListByRepoOptions{
		Labels: []string{
			defectLabel,
			teamLabel,
		},
		Since: lastMonth,
	}

	lastMonthDefects, err := issues.GetRepoIssues(owner, repo, options)
	if err != nil {
		return -1, nil
	}
	return len(lastMonthDefects), nil

}
