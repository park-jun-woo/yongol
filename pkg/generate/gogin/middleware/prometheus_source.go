//ff:type feature=gen-gogin type=generator
//ff:what prometheusSourceTemplate — prometheus.go 를 2파일로 분할하는 소스 템플릿 맵

package middleware

// prometheusMiddlewareSourceTemplate carries the source for
// internal/middleware/prometheus_middleware.go.
// __BUCKETS__ is replaced with the concrete buckets literal.
const prometheusMiddlewareSourceTemplate = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=observability
//` + `ff:what PrometheusMiddleware — /metrics 기본 3종 지표 자동 수집 미들웨어

package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// httpRequestsTotal counts completed HTTP requests labelled by method, path
// and status. path is always the gin route template (c.FullPath()) so label
// cardinality stays bounded.
var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests processed, labelled by method, path and status.",
	},
	[]string{"method", "path", "status"},
)

// httpRequestDuration observes request latency in seconds labelled by
// method, path and status.
var httpRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "HTTP request latency in seconds, labelled by method, path and status.",
		Buckets: __BUCKETS__,
	},
	[]string{"method", "path", "status"},
)

// httpRequestsInFlight gauges the number of currently in-flight requests per
// (method, path). Decremented on handler return (deferred).
var httpRequestsInFlight = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "http_requests_in_flight",
		Help: "Number of in-flight HTTP requests, labelled by method and path.",
	},
	[]string{"method", "path"},
)

func init() {
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration, httpRequestsInFlight)
}

// PrometheusMiddleware returns a gin middleware that records the three core
// HTTP metrics. Requests that do not match any route (c.FullPath() == "")
// are bucketed under a literal "unmatched" path label so cardinality is
// capped even for 404 scan traffic.
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.FullPath()
		if path == "" {
			path = "unmatched"
		}
		httpRequestsInFlight.WithLabelValues(method, path).Inc()
		defer httpRequestsInFlight.WithLabelValues(method, path).Dec()

		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		httpRequestsTotal.WithLabelValues(method, path, status).Inc()
		httpRequestDuration.WithLabelValues(method, path, status).Observe(time.Since(start).Seconds())
	}
}
`

// prometheusHandlerSource carries the source for
// internal/middleware/prometheus_handler.go.
const prometheusHandlerSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=observability
//` + `ff:what PrometheusHandler — Prometheus exposition format 으로 /metrics 핸들러 제공

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PrometheusHandler returns a gin handler that exposes the Prometheus
// default registry in the text exposition format. Mounted on the scrape
// path by generated main.go.
func PrometheusHandler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}
`
