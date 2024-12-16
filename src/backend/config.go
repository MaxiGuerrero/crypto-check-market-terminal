package backend

import (
	"os"
	"strconv"
)

type Config struct {
	token          string
	amountCurrency int
}

func NewConfig() *Config {
	return &Config{
		token:          os.Getenv("TOKEN"),
		amountCurrency: getAmountCurrencyEnv(),
	}
}

func getAmountCurrencyEnv() int {
	data := os.Getenv("AMOUNT_CURRENCIES")
	amountCurrencies, _ := strconv.Atoi(data)
	return amountCurrencies
}
