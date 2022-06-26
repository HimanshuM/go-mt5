package mt5tests

import (
	"os"
	"testing"

	"github.com/HimanshuM/go_mt5/mt5"
	"github.com/stretchr/testify/require"
)

func testLogin(t *testing.T) {
	mt = &mt5.MT5{}
	err := mt.Init(&mt5.MT5Config{
		Host:        os.Getenv("MT5_HOST"),
		Port:        os.Getenv("MT5_PORT"),
		Username:    os.Getenv("MT5_USERNAME"),
		Password:    os.Getenv("MT5_PASSWORD"),
		CryptMethod: "NONE",
	})
	require.NoErrorf(t, err, "error during login: %v", err)
}
