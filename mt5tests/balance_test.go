package mt5tests

import (
	"os"
	"testing"

	"github.com/HimanshuM/go_mt5/mt5"
	"github.com/stretchr/testify/require"
)

var deposit = 100

func testBalanceDeposit(t *testing.T) {
	trade := &mt5.Trade{
		Login:       os.Getenv("USER_LOGIN"),
		Amount:      int64(deposit),
		Comment:     "Deposit test Go wrapper",
		CheckMargin: true,
	}
	err := mt.SetBalance(trade)
	require.NoErrorf(t, err, "error during updating balance: %v", err)
	require.NotEmpty(t, trade.Ticket)
}

func testBalanceWithdrawSuccess(t *testing.T) {
	trade := &mt5.Trade{
		Login:       os.Getenv("USER_LOGIN"),
		Amount:      -50,
		Comment:     "Withdraw test Go wrapper",
		CheckMargin: true,
	}
	err := mt.SetBalance(trade)
	require.NoErrorf(t, err, "error during updating balance: %v", err)
	require.NotEmpty(t, trade.Ticket)
}

func testBalanceWithdrawFail(t *testing.T) {
	trade := &mt5.Trade{
		Login:       os.Getenv("USER_LOGIN"),
		Amount:      -150,
		Comment:     "Withdraw test Go wrapper fail",
		CheckMargin: true,
	}
	err := mt.SetBalance(trade)
	require.Error(t, err)
	require.EqualError(t, err, "error setting balance: No money")
	require.Empty(t, trade.Ticket)
}

func testBalanceWithdrawNoCheck(t *testing.T) {
	trade := &mt5.Trade{
		Login:       os.Getenv("USER_LOGIN"),
		Amount:      -150,
		Comment:     "Withdraw test Go wrapper no check",
		CheckMargin: false,
	}
	err := mt.SetBalance(trade)
	require.NoErrorf(t, err, "error during updating balance: %v", err)
	require.NotEmpty(t, trade.Ticket)
	deposit = 500
	t.Run("testBalanceDepositRestore", testBalanceDeposit)
}
