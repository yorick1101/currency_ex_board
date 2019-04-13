package main

import (
	"log"
	"sync"
	"time"

	"github.com/yorick1101/currency_ex_board/internal/dao"
	service "github.com/yorick1101/currency_ex_board/internal/service"
)

var currencyDao = dao.NewCurrencyDao()
var wg = new(sync.WaitGroup)

func main() {
	log.Println("Start exchange server")
	wg.Add(1)
	duration, err := time.ParseDuration("10s")
	if err != nil {
		log.Fatal("failed to parse duration")
	}

	task := service.NewTask(crawlTask, duration)

	go func() {
		task.Run()
		wg.Done()
	}()
	wg.Wait()
}

func crawlTask() error {
	newRates := service.Crawl()
	log.Println("scheduler fired", time.Now().Format("2006-01-02 15:04:05"))
	for _, rate := range newRates {
		currencyDao.AddExchangeRate(rate)
	}
	return nil
}
