package promfasthttp_test

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/trim21/promfasthttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func BenchmarkPromFasthttp(b *testing.B) {
	ctx := &fasthttp.RequestCtx{
		Request:  fasthttp.Request{Header: fasthttp.RequestHeader{}},
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
		Request:  fasthttp.Request{Header: fasthttp.RequestHeader{}},
		Response: fasthttp.Response{},
	}
	reg := prometheus.NewRegistry()

	handler := fasthttpadaptor.NewFastHTTPHandler(promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	for i := 0; i < b.N; i++ {
		handler(ctx)
	}
}
