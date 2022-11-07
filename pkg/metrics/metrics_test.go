package metrics

import (
	"github.com/enekofb/metrics/pkg/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateQueryFuncFromConfig(t *testing.T) {

	config := config.Config{
		GithubConfig: config.GithubConfig{
			Token:   "abc",
			Queries: nil,
		},
		MetricsConfig: []config.MetricConfig{
			{
				Name: "metric_test",
				Type: "gauge",
			},
		},
	}

	metrics := CreateMetricsFromConfig(config)
	for _, metric := range metrics {
		require.Equal(t, metric.Name, "metric_test")

	}

}
