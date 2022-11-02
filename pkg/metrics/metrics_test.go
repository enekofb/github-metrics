package metrics

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetLastMonthDefectMetricsByTeam(t *testing.T) {

	numDefectsByTeam, err := GetLastMonthDefectMetricsByTeam()
	require.NoError(t, err)
	require.NotEmpty(t, numDefectsByTeam > 0)
}
