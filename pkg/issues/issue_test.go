package issues

import (
	github2 "github.com/google/go-github/v48/github"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCanGetDefectsForRepository(t *testing.T) {
	owner := os.Getenv("GITHUB_OWNER")
	require.NotEmpty(t, owner)

	repo := os.Getenv("GITHUB_REPO")
	require.NotEmpty(t, repo)

	defectLabel := os.Getenv("GITHUB_BUG_LABEL")
	require.NotEmpty(t, defectLabel)

	teamLabel := os.Getenv("GITHUB_TEAM_LABEL")
	require.NotEmpty(t, teamLabel)

	options := &github2.IssueListByRepoOptions{
		Labels: []string{
			defectLabel,
			teamLabel,
		},
	}
	defects, err := GetRepoIssues(owner, repo, options)
	require.NoError(t, err)
	require.NotEmpty(t, defects)
}

func TestGetRepoIssues(t *testing.T) {
	owner := os.Getenv("GITHUB_OWNER")
	require.NotEmpty(t, owner)

	repo := os.Getenv("GITHUB_REPO")
	require.NotEmpty(t, repo)

	rep, err := GetRepoIssues(owner, repo, &github2.IssueListByRepoOptions{})
	require.NoError(t, err)
	require.NotEmpty(t, rep)
}
