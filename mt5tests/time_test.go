package mt5tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testTimestampGet(t *testing.T) {
	timestamp, err := mt.GetServerTime()
	require.NoError(t, err)
	require.NotEmpty(t, timestamp)
}
