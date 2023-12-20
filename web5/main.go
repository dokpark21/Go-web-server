package main

import (
	"log"
	"net/http"
	"time"

	"example5.com/decohandler"
	"example5.com/myapp"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed time:", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed time:", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	h := myapp.NewHandler()
	// logger1이 httpHandler를 감싸고 있다.
	h = decohandler.NewDecoHandler(h, logger)
	// logger2가 httpHandle를 감싸고 있다. 순서 : logger2 Start -> logger1 Start -> httpHandler -> logger1 Completed -> logger2 Completed
	h = decohandler.NewDecoHandler(h, logger2)
	return h
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}
