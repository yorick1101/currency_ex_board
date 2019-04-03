package dao

import (
	"log"
	"strconv"
	"time"
)

type CurrencyDao interface {
	GetExchangeRates(quote string, base string) []*ExchangeRate
	AddExchangeRate(rate *ExchangeRate)
}

type ExchangeRate struct {
	Quote string
	Base  string
	Date  *time.Time
	Rate  float64
}

func BuildExchangeRate(quote string, base string, exchange_rate_str string, date *time.Time) *ExchangeRate {
	exchange_rate, err := strconv.ParseFloat(exchange_rate_str, 64)
	if err != nil {
		log.Fatal("failed to parse exchange rate", exchange_rate_str, err)
	}
	rate := new(ExchangeRate)
	rate.Base = base
	rate.Quote = quote
	rate.Date = date
	rate.Rate = exchange_rate
	return rate
}
