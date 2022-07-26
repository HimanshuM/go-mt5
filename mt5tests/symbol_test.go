package mt5tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testSymbolGet(t *testing.T) {
	symbol, err := mt.GetSymbol("EURUSD")
	require.NoError(t, err)
	require.NotEmpty(t, symbol)
}

func testSymbolsGetAll(t *testing.T) {
	symbols, err := mt.GetAllSymbols()
	require.NoError(t, err)
	require.NotEmpty(t, symbols)
}

func testSymbolsGetByIndex(t *testing.T) {
	symbol, err := mt.GetSymbolByIndex(0)
	require.NoError(t, err)
	require.NotEmpty(t, symbol)
}
