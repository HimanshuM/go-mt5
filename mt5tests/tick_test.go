package mt5tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testGetLastTick(t *testing.T) {
	lastTick, err := mt.GetLastTick(0, "EURUSD")
	require.NoError(t, err)
	require.NotEmpty(t, lastTick)
	require.NotEmpty(t, lastTick.TransactionID)
	require.Equal(t, 1, len(lastTick.Ticks))
}

func testGetLastTicksMultiple(t *testing.T) {
	lastTicks, err := mt.GetLastTick(0, "EURUSD", "GBPUSD")
	require.NoError(t, err)
	require.NotEmpty(t, lastTicks)
	require.NotEmpty(t, lastTicks.TransactionID)
	require.Equal(t, 2, len(lastTicks.Ticks))
}

func testGetTickStats(t *testing.T) {
	lastTicks, err := mt.GetTickStatistics(0, "EURUSD", "GBPUSD")
	require.NoError(t, err)
	require.NotEmpty(t, lastTicks)
	require.NotEmpty(t, lastTicks.TransactionID)
	require.Equal(t, 2, len(lastTicks.TickStats))
}
