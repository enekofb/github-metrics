package github

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClientFromEnvironment()
	require.NoError(t, err)
	require.NotNil(t, client)
}
