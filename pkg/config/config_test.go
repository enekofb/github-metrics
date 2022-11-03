package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRead(t *testing.T) {
	metricsConfig, err := Read("../../resources")
	require.NoError(t, err)
	githubToken := metricsConfig.githubConfig.token
	require.Equal(t, githubToken, "abc")
}
