package main

import (
	"fmt"
	"os"

	"github.com/HimanshuM/go_mt5/mt5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Please define a `.env` file with the required config before proceeding.")
		return
	}

	mt := mt5.MT5{}
	err := mt.Init(&mt5.MT5Config{
		Host:        os.Getenv("MT5_HOST"),
		Port:        os.Getenv("MT5_PORT"),
		Username:    os.Getenv("MT5_USERNAME"),
		Password:    os.Getenv("MT5_PASSWORD"),
		CryptMethod: "NONE",
	})
	if err != nil {
		logrus.Errorf("error during login: %v", err)
		return
	}
	trade := &mt5.Trade{
		Login:       os.Getenv("USER_LOGIN"),
		Amount:      10,
		Comment:     "Deposit test Go wrapper",
		CheckMargin: true,
	}
	err = mt.SetBalance(trade)
	if err != nil {
		logrus.Errorf("error during updating balance: %v", err)
	} else {
		logrus.Infof("Ticket number: %v", trade.Ticket)
	}

	trade = &mt5.Trade{
		Login:       os.Getenv("USER_LOGIN"),
		Amount:      -1000000,
		Comment:     "Withdraw test Go wrapper",
		CheckMargin: true,
	}
	err = mt.SetBalance(trade)
	if err != nil {
		logrus.Errorf("error during updating balance: %v", err)
	} else {
		logrus.Infof("Ticket number: %v", trade.Ticket)
	}
}
