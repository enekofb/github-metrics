package issues

import (
	"context"
	"github.com/enekofb/metrics/pkg/github"
	github2 "github.com/google/go-github/v48/github"
)

func GetRepoIssues(owner string, repo string, options *github2.IssueListByRepoOptions) ([]*github2.Issue, error) {

	c, err := github.NewClientFromEnvironment()
	if err != nil {
		return nil, err
	}

	//TODO validate options

	issues, _, err := c.Issues.ListByRepo(context.Background(), owner, repo, options)
	if err != nil {
		return nil, err
	}

	return issues, nil

}
