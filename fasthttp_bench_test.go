package promfasthttp_test

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/trim21/promfasthttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func requestHeader() *fasthttp.RequestHeader {
	raw := map[string]string{
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"accept-language":           "en-US,en",
		"cache-control":             "no-cache",
		"pragma":                    "no-cache",
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "none",
		"sec-fetch-user":            "?1",
		"sec-gpc":                   "1",
		"upgrade-insecure-requests": "1",
	}

	h := &fasthttp.RequestHeader{}

	for key, value := range raw {
		h.Set(key, value)
	}

	return h
}

func BenchmarkPromFasthttp(b *testing.B) {
	ctx := &fasthttp.RequestCtx{
		Request:  fasthttp.Request{Header: *requestHeader()},
		Response: fasthttp.Response{},
	}
	reg := prometheus.NewRegistry()

	handler := promfasthttp.HandlerFor(reg, promfasthttp.HandlerOpts{})

	for i := 0; i < b.N; i++ {
		handler(ctx)
	}
}

func BenchmarkPromHTTP(b *testing.B) {
	ctx := &fasthttp.RequestCtx{
		Request:  fasthttp.Request{Header: *requestHeader()},
		Response: fasthttp.Response{},
	}
	reg := prometheus.NewRegistry()

	handler := fasthttpadaptor.NewFastHTTPHandler(promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	for i := 0; i < b.N; i++ {
		handler(ctx)
	}
}
