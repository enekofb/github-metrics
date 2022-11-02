package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
	"os"
)

func NewClientFromEnvironment() (*github.Client, error) {

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("github token not found")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client, nil

}
