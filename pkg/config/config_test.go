package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRead(t *testing.T) {
	metricsConfig, err := Read("../../resources/test/config.yaml")
	require.NoError(t, err)
	githubToken := metricsConfig.GithubConfig.Token
	require.Equal(t, githubToken, "abc")
}
