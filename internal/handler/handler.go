package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	dao "github.com/yorick1101/currency_ex_board/internal/dao"
)

var currencyDao = dao.NewCurrencyDao()

type HttpHandler struct {
}

func NewHandler() http.Handler {

	return new(HttpHandler)
}

func (*HttpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	parsedurl, err := url.ParseRequestURI(request.RequestURI)
	if err != nil {
		log.Panic("failed to parse url")
		return
	}

	values, err := url.ParseQuery(parsedurl.RawQuery)
	if err != nil {
		log.Panic("failed to parser query string")
		return
	}

	base := values.Get("base")
	quote := values.Get("quote")
	if base == "" || quote == "" {
		log.Panic("lack of query parameter base or quote")
		return
	}

	rates := currencyDao.GetExchangeRates(base, quote)
	slcB, _ := json.Marshal(rates)
	response.Write(slcB)
}
