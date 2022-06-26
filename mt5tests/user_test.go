package mt5tests

import (
	"testing"

	"github.com/HimanshuM/go_mt5/mt5"
	"github.com/stretchr/testify/require"
)

func testUserCreate(t *testing.T) {
	user := &mt5.MT5User{
		Name:           "Go Test",
		Email:          "go@test.com",
		Rights:         0x1E3,
		Group:          "demo\\forex",
		Leverage:       100,
		MainPassword:   "QWEasdZXD",
		InvestPassword: "QWEasdZXD",
		Color:          0xFF000000,
	}
	err := mt.CreateUser(user)
	require.NoErrorf(t, err, "error creating user: %v", err)
	require.NotEmpty(t, user.Login)
}
