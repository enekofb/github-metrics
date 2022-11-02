package metrics

import (
	"github.com/enekofb/metrics/pkg/issues"
	github2 "github.com/google/go-github/v48/github"
	"os"
	"time"
)

func GetLastMonthDefectMetricsByTeam() (int, error) {

	//TODO we want this in config map
	owner := os.Getenv("GITHUB_OWNER")
	if owner == "" {
		return -1, nil
	}

	repo := os.Getenv("GITHUB_REPO")
	if repo == "" {
		return -1, nil
	}

	defectLabel := os.Getenv("GITHUB_BUG_LABEL")
	if defectLabel == "" {
		return -1, nil
	}

	teamLabel := os.Getenv("GITHUB_TEAM_LABEL")
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
