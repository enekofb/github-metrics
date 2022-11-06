package metrics

import (
	"github.com/enekofb/metrics/pkg/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateQueryFuncFromConfig(t *testing.T) {

	var queryFunc QueryFunc
	queryConfig := config.QueryConfig{
		Owner:     "kubernetes",
		Repo:      "kubernetes",
		BugLabel:  "kind/bug",
		TeamLabel: "sig/apps",
	}
	queryFunc = NewFromConfig(queryConfig)
	require.NotNil(t, queryFunc)
	i, err := queryFunc()
	require.NoError(t, err)
	require.True(t, i > 0)

}
