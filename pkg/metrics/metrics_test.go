package metrics

import (
	"github.com/enekofb/metrics/pkg/config"
	"github.com/stretchr/testify/require"
	"testing"
)

// TODO: test me
func TestCreateQueryFuncFromConfig(t *testing.T) {

	var queryFuncs map[string]QueryFunc
	queryConfig := config.QueryConfig{
		Owner:     "kubernetes",
		Repo:      "kubernetes",
		BugLabel:  "kind/bug",
		TeamLabel: "sig/apps",
	}

	queryFuncs = CreateMetricsFromConfig([]config.QueryConfig{queryConfig})
	require.NotEmpty(t, queryFuncs)
	for queryName, queryFunc := range queryFuncs {
		require.Equal(t, queryName, queryConfig.Name)
		i, err := queryFunc()
		require.NoError(t, err)
		require.True(t, i > 0)
	}

}
