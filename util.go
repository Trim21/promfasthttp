package promfasthttp

import (
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"
)

// fasthttp version of http.Error()
func stdFastHTTPError(ctx *fasthttp.RequestCtx, error string, code int) {
	ctx.SetStatusCode(code)
	ctx.Response.Header.Set("Content-Type", "text/plain; charset=utf-8")
	ctx.Response.Header.Set("X-Content-Type-Options", "nosniff")
	fmt.Fprintln(ctx, error)
}

func httpError(ctx *fasthttp.RequestCtx, err error) {
	ctx.Response.Header.Del(fasthttp.HeaderContentEncoding)
	stdFastHTTPError(
		ctx,
		"An error has occurred while serving metrics:\n\n"+err.Error(),
		http.StatusInternalServerError,
	)
}
