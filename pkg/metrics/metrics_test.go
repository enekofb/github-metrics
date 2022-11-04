package metrics

import (
	"github.com/enekofb/metrics/pkg/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetLastMonthDefectMetricsByTeam(t *testing.T) {

	//TODO complete me
	config := config.MetricsConfig{}

	numDefectsByTeam, err := GetLastMonthDefectMetricsByTeam(config)
	require.NoError(t, err)
	require.NotEmpty(t, numDefectsByTeam > 0)
}
