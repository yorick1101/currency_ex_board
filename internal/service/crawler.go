package service

import (
	"encoding/csv"
	"log"
	"net/http"
	"time"

	dao "github.com/yorick1101/currency_ex_board/internal/dao"
)

func Crawl() []*dao.ExchangeRate {
	resp, err := http.Get("http://www.taifex.com.tw/data_gov/taifex_open_data.asp?data_name=DailyForeignExchangeRates")
	if err != nil {
		log.Fatal("no respnse", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		records, err := csv.NewReader(resp.Body).ReadAll()
		if err != nil {
			log.Fatal("failed to read csv from http response")
		}
		//remove header
		records = records[1:]
		rates := make([]*dao.ExchangeRate, 0)

		for _, row := range records {
			create_date, err := time.Parse("2006/1/2", row[0])
			if err != nil {
				log.Fatal("failed to parse date string", err)
			}
			rates = append(rates, dao.BuildExchangeRate("USD", "NTD", row[1], &create_date))
			rates = append(rates, dao.BuildExchangeRate("CNY", "NTD", row[2], &create_date))
			rates = append(rates, dao.BuildExchangeRate("EUR", "USD", row[3], &create_date))
			rates = append(rates, dao.BuildExchangeRate("USD", "JPY", row[4], &create_date))
			rates = append(rates, dao.BuildExchangeRate("GBP", "USD", row[5], &create_date))
			rates = append(rates, dao.BuildExchangeRate("AUD", "USD", row[6], &create_date))
			rates = append(rates, dao.BuildExchangeRate("USD", "HKD", row[7], &create_date))
			rates = append(rates, dao.BuildExchangeRate("USD", "CNY", row[8], &create_date))
			rates = append(rates, dao.BuildExchangeRate("USD", "ZAR", row[9], &create_date))
			rates = append(rates, dao.BuildExchangeRate("NZD", "USD", row[10], &create_date))
		}

		return rates
	}

	return nil
}
