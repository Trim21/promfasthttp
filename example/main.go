package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/trim21/promfasthttp"
	"github.com/valyala/fasthttp"
)

func main() {
	go http.ListenAndServe("127.0.0.1:8091", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))

	h := promfasthttp.HandlerFor(prometheus.DefaultGatherer, promfasthttp.HandlerOpts{})

	if err := fasthttp.ListenAndServe("127.0.0.1:8089", h); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
