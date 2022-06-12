# PromFasthttp

fasthttp request handler for prometheus.

`fasthttpadaptor.NewFastHTTPHandler()` works fine,
and there is no significant improvement (see [#benchmark]), you may not need this.

## example

```golang
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
```

## benchmark

comparing `promfasthttp.HandlerFor(reg, promfasthttp.HandlerOpts{})` with
`fasthttpadaptor.NewFastHTTPHandler(promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))`

```text
goos: windows
goarch: amd64
pkg: github.com/trim21/promfasthttp
cpu: AMD Ryzen 7 5800X 8-Core Processor
BenchmarkPromFasthttp-16          154302              7688 ns/op           33525 B/op         19 allocs/op
BenchmarkPromHTTP-16              137118              8912 ns/op           34640 B/op         32 allocs/op
PASS
ok      github.com/trim21/promfasthttp  2.817s
```

## difference

HandlerOpts.DisableCompression is removed,
you should use [`fasthttp.CompressHandler`](https://pkg.go.dev/github.com/valyala/fasthttp#CompressHandler) instead.

```golang
package main

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/trim21/promfasthttp"
	"github.com/valyala/fasthttp"
)

func main() {
	h := fasthttp.CompressHandler(promfasthttp.HandlerFor(prometheus.DefaultGatherer, promfasthttp.HandlerOpts{}))

	if err := fasthttp.ListenAndServe(":8080", h); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
```
