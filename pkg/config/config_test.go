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

func TestCanReadQueryConfigurations(t *testing.T) {
	metricsConfig, err := Read("../../resources/test/configWithQueries.yaml")
	require.NoError(t, err)
	queries := metricsConfig.GithubConfig.Queries
	require.NotEmpty(t, queries)
	query := queries[0]
	require.Equal(t, query.Name, "basic_query")
	require.Equal(t, query.BugLabel, "bug")
}

func TestCanReadMetricConfigurations(t *testing.T) {
	config, err := Read("../../resources/test/configWithMetrics.yaml")
	require.NoError(t, err)
	metrics := config.MetricsConfig
	require.NotEmpty(t, metrics)
	metric := metrics[0]
	require.Equal(t, metric.Name, "defects")
	require.Equal(t, metric.Type, "gauge")
}
