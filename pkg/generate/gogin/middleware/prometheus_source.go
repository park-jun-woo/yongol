//ff:type feature=gen-gogin type=generator
//ff:what prometheusSourceTemplate — GeneratePrometheus 가 기록하는 prometheus.go 템플릿

package middleware

import (
	"fmt"
	"strings"
)

// prometheusSourceTemplate carries the verbatim Go source for
// internal/middleware/prometheus.go. __BUCKETS__ is replaced with the
// concrete buckets literal from the manifest (default prometheus.DefBuckets).
// Provides:
//
//   - PrometheusMiddleware()     — gin.HandlerFunc measuring requests.
//   - PrometheusHandler()        — gin handler wrapping promhttp.Handler.
//   - http_requests_total        — Counter (method, path, status).
//   - http_request_duration_seconds — Histogram (method, path, status).
//   - http_requests_in_flight    — Gauge (method, path).
//
// path labels use c.FullPath() so the label set is the OpenAPI route
// template (e.g. "/users/{id}"), not the concrete value — keeps cardinality
// bounded.
const prometheusSourceTemplate = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=observability
//` + `ff:what PrometheusMiddleware / PrometheusHandler — /metrics + 기본 3종 지표 자동 수집

package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

// PrometheusHandler returns a gin handler that exposes the Prometheus
// default registry in the text exposition format. Mounted on the scrape
// path by generated main.go.
func PrometheusHandler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}
`

// renderPrometheusSource substitutes __BUCKETS__ with the buckets literal.
// When buckets is empty prometheus.DefBuckets is emitted so operators can
// still tune the slice without regenerating.
func renderPrometheusSource(buckets []float64) string {
	lit := bucketsLiteral(buckets)
	return strings.ReplaceAll(prometheusSourceTemplate, "__BUCKETS__", lit)
}

// bucketsLiteral formats the histogram buckets as a Go expression. An empty
// or nil slice yields "prometheus.DefBuckets" so the generated code keeps
// the widely-recognised web-API default (0.005s ~ 10s).
func bucketsLiteral(buckets []float64) string {
	if len(buckets) == 0 {
		return "prometheus.DefBuckets"
	}
	parts := make([]string, 0, len(buckets))
	for _, b := range buckets {
		parts = append(parts, fmt.Sprintf("%g", b))
	}
	return "[]float64{" + strings.Join(parts, ", ") + "}"
}
