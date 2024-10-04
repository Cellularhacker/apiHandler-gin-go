package fasthttp

import (
	"github.com/Cellularhacker/apiError-go"
	"github.com/valyala/fasthttp"
)

type FastHttp func(*fasthttp.RequestCtx) *apiError.Error

func (fhh FastHttp) Handler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

	}
}
