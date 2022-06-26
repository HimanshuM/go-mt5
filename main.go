package main

import (
	"fmt"
	"os"

	"github.com/HimanshuM/go_mt5/mt5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var mt *mt5.MT5

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Please define a `.env` file with the required config before proceeding.")
		return
	}

	if err := login(); err != nil {
		return
	}

	if err := balance_success(); err != nil {
		return
	}

	if err := balance_fail(); err == nil {
		return
	}

	if err := user_create(); err != nil {
		return
	}

}

func login() error {
	mt = &mt5.MT5{}
	err := mt.Init(&mt5.MT5Config{
		Host:        os.Getenv("MT5_HOST"),
		Port:        os.Getenv("MT5_PORT"),
		Username:    os.Getenv("MT5_USERNAME"),
		Password:    os.Getenv("MT5_PASSWORD"),
		CryptMethod: "NONE",
	})
	if err != nil {
		logrus.Errorf("error during login: %v", err)
		return err
	}
	return err
}

func balance_success() error {
	trade := &mt5.Trade{
		Login:       os.Getenv("USER_LOGIN"),
		Amount:      10,
		Comment:     "Deposit test Go wrapper",
		CheckMargin: true,
	}
	err := mt.SetBalance(trade)
	if err != nil {
		logrus.Errorf("error during updating balance: %v", err)
	} else {
		logrus.Infof("Ticket number: %v", trade.Ticket)
	}
	return err
}

func balance_fail() error {
	trade := &mt5.Trade{
		Login:       os.Getenv("USER_LOGIN"),
		Amount:      -1000000,
		Comment:     "Withdraw test Go wrapper",
		CheckMargin: true,
	}
	err := mt.SetBalance(trade)
	if err != nil {
		logrus.Errorf("error during updating balance: %v", err)
	} else {
		logrus.Infof("Ticket number: %v", trade.Ticket)
	}
	return err
}

func user_create() error {
	user := &mt5.MT5User{
		Name:           "Go Test",
		Email:          "go@test.com",
		Rights:         0x1E3,
		Group:          "demo\\forex",
		Leverage:       100,
		MainPassword:   "QWEasdZXC",
		InvestPassword: "QWEasdZXC",
		Color:          0xFF000000,
	}
	err := mt.CreateUser(user)
	if err != nil {
		logrus.Errorf("error creating user: %v", err)
	} else {
		logrus.Infof("user create: %+v", user)
	}
	return err
}
