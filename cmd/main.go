package main

import (
	"log"

	dao "github.com/yorick1101/currency_ex_board/internal/dao"
	crawler "github.com/yorick1101/currency_ex_board/internal/service"
)

func main() {
	log.Println("Start exchange server")
	currencyDao := dao.NewCurrencyDao()

	newRates := crawler.Crawl()

	for _, rate := range newRates {
		currencyDao.AddExchangeRate(rate)
	}

	currencyDao.GetExchangeRates("NZD", "USD")
}
