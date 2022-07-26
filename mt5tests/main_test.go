package mt5tests

import (
	"testing"

	"github.com/HimanshuM/go-mt5/mt5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var mt *mt5.MT5

func TestMain(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	err := godotenv.Load("../.env")
	require.NoError(t, err, "Please define a `.env` file with the required config before proceeding.")

	t.Run("testLogin", testLogin)

	// Balance tests
	t.Run("testBalance", testBalanceDeposit)
	t.Run("testBalanceWithdrawSuccess", testBalanceWithdrawSuccess)
	t.Run("testBalanceWithdrawFail", testBalanceWithdrawFail)
	t.Run("testBalanceWithdrawNoCheck", testBalanceWithdrawNoCheck)

	// User tests
	t.Run("testUserCreate", testUserCreate)

	// Symbol tests
	t.Run("testSymbolGet", testSymbolGet)
	t.Run("testSymbolsGetAll", testSymbolsGetAll)
	t.Run("testSymbolsGetByIndex", testSymbolsGetByIndex)

	// Time tests
	t.Run("testTimestampGet", testTimestampGet)
}
