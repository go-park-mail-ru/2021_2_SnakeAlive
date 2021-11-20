package http

import "github.com/valyala/fasthttp"

type Middleware func(fasthttp.RequestHandler) fasthttp.RequestHandler
