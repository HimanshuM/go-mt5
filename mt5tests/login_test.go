package mt5tests

import (
	"os"
	"testing"

	"github.com/HimanshuM/go-mt5/mt5"
	"github.com/stretchr/testify/require"
)

func testLogin(t *testing.T) {
	mt = &mt5.Client{}
	err := mt.Init(&mt5.Config{
		Host:        os.Getenv("MT5_HOST"),
		Port:        os.Getenv("MT5_PORT"),
		Username:    os.Getenv("MT5_USERNAME"),
		Password:    os.Getenv("MT5_PASSWORD"),
		CryptMethod: "NONE",
	})
	require.NoErrorf(t, err, "error during login: %v", err)
}
