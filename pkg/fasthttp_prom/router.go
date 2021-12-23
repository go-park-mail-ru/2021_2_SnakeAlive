package fasthttpprom

import (
	"net/http"
	"strconv"
	"time"

	fst_http "github.com/fasthttp/router"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var defaultMetricPath = "/metrics"

type Router interface {
	GET(path string, handler fasthttp.RequestHandler)
	POST(path string, handler fasthttp.RequestHandler)
	PATCH(path string, handler fasthttp.RequestHandler)
	DELETE(path string, handler fasthttp.RequestHandler)
	PUT(path string, handler fasthttp.RequestHandler)

	Use(r *fst_http.Router)

	GetHandler() fasthttp.RequestHandler
}

type router struct {
	r          *fst_http.Router
	subsystem  string
	metricPath string
	reqDur     *prometheus.HistogramVec
}

func (r *router) GET(path string, handler fasthttp.RequestHandler) {
	r.r.GET(path, r.metricMw(handler, path, http.MethodGet))
}

func (r *router) POST(path string, handler fasthttp.RequestHandler) {
	r.r.POST(path, r.metricMw(handler, path, http.MethodPost))
}

func (r *router) PATCH(path string, handler fasthttp.RequestHandler) {
	r.r.PATCH(path, r.metricMw(handler, path, http.MethodPatch))
}

func (r *router) DELETE(path string, handler fasthttp.RequestHandler) {
	r.r.DELETE(path, r.metricMw(handler, path, http.MethodDelete))
}

func (r *router) PUT(path string, handler fasthttp.RequestHandler) {
	r.r.PUT(path, r.metricMw(handler, path, http.MethodPut))
}

func (r *router) GetHandler() fasthttp.RequestHandler {
	return r.r.Handler
}

func (r *router) Use(fR *fst_http.Router) {
	r.r = fR
	fR.GET(r.metricPath, prometheusHandler())
}

func (r *router) metricMw(next fasthttp.RequestHandler, path, method string) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()

		next(ctx)

		if string(ctx.Request.URI().Path()) == "/favicon.ico" {
			return
		}

		status := strconv.Itoa(ctx.Response.StatusCode())
		elapsed := float64(time.Since(start)) / float64(time.Second)
		r.reqDur.WithLabelValues(status, path, method).Observe(elapsed)

	}
}

func NewRouter(subsystem string) Router {
	reqDur := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: subsystem,
			Name:      "request_duration_seconds",
			Help:      "request latencies",
			Buckets:   []float64{.005, .01, .02, 0.04, .06, 0.08, .1, 0.15, .25, 0.4, .6, .8, 1, 1.5, 2, 3, 5},
		},
		[]string{"code", "path", "method"},
	)
	_ = prometheus.Register(reqDur)

	return &router{
		subsystem:  subsystem,
		metricPath: defaultMetricPath,
		reqDur:     reqDur,
	}
}

// since prometheus/client_golang use net/http we need this net/http adapter for fasthttp
func prometheusHandler() fasthttp.RequestHandler {
	return fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
}
