package main

import (
	"log"
	"net/http"
	"time"

	"github.com/yorick1101/currency_ex_board/internal/handler"
)

func main() {

	myHandler := handler.NewHandler()
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.Handle("/rate", myHandler)
	log.Fatal(s.ListenAndServe())
}
