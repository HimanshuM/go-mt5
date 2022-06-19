package main

import (
	"fmt"
	"os"

	"github.com/HimanshuM/go_mt5/mt5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Please define a `.env` file with the required config before proceeding.")
		return
	}

	mt := mt5.MT5{}
	mt.Init(&mt5.MT5Config{
		Host:     os.Getenv("MT5_HOST"),
		Port:     os.Getenv("MT5_PORT"),
		Username: os.Getenv("MT5_USERNAME"),
		Password: os.Getenv("MT5_PASSWORD"),
	})
}
