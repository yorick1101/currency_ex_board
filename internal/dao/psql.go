package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewCurrencyDao() CurrencyDao {

	/*
		os.Setenv("db-username", "yorick")
		os.Setenv("db-password", "yorick123")
		os.Setenv("db-url", "localhost")
		os.Setenv("db-port", "5432")
		os.Setenv("db-database", "currency")
	*/
	username := os.Getenv("db-username")
	password := os.Getenv("db-password")
	url := os.Getenv("db-url")
	port := os.Getenv("db-port")
	database := os.Getenv("db-database")
	connection_str := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, url, port, database)
	log.Println("connect postgres:", connection_str)
	dao := new(PsqlCurrencyDao)
	db, err := sql.Open("postgres", connection_str)
	if err != nil {
		log.Fatal("failed to open db connection")
	}
	dao.db = db
	return dao
}

type PsqlCurrencyDao struct {
	db *sql.DB
}

func (dao *PsqlCurrencyDao) GetExchangeRates(quote string, base string) []*ExchangeRate {
	rows, err := dao.db.Query("select * from exchange_rate where base=$1 and quote=$2", base, quote)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var results []*ExchangeRate
	for rows.Next() {
		var (
			id             int64
			base_currency  string
			quote_currency string
			exchange_rate  float64
			create_date    time.Time
		)
		if err := rows.Scan(&id, &base_currency, &quote_currency, &exchange_rate, &create_date); err != nil {
			log.Fatal(err)
		}
		result := new(ExchangeRate)
		result.Base = base_currency
		result.Quote = quote_currency
		result.Date = &create_date
		result.Rate = exchange_rate
		results = append(results, result)
		log.Print("base:", base, ";quote:", quote, ";rate:", exchange_rate, ";date:", create_date)
	}
	log.Print("size", len(results))
	return results
}

func (dao *PsqlCurrencyDao) AddExchangeRate(rate *ExchangeRate) {
	stmt, err := dao.db.Prepare("insert into exchange_rate(base, quote, rate, date) values($1, $2, $3, $4) on conflict on constraint unique_exchange_rate do nothing ")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(rate.Base, rate.Quote, rate.Rate, rate.Date)
}
