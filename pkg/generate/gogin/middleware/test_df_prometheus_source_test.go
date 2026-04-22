//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what prometheus_source 렌더 스냅샷 — buckets 기본/커스텀

package middleware

import (
	"strings"
	"testing"
)

func TestRenderPrometheusSource_DefaultBuckets(t *testing.T) {
	out := renderPrometheusSource(nil)
	if !strings.Contains(out, "prometheus.DefBuckets") {
		t.Fatalf("expected prometheus.DefBuckets literal for nil buckets, got:\n%s", out)
	}
	for _, must := range []string{
		"func PrometheusMiddleware()",
		"func PrometheusHandler()",
		`"http_requests_total"`,
		`"http_request_duration_seconds"`,
		`"http_requests_in_flight"`,
		"promhttp.Handler()",
		"c.FullPath()",
	} {
		if !strings.Contains(out, must) {
			t.Errorf("rendered source missing fragment %q", must)
		}
	}
}

func TestRenderPrometheusSource_CustomBuckets(t *testing.T) {
	out := renderPrometheusSource([]float64{0.1, 0.5, 1})
	if !strings.Contains(out, "[]float64{0.1, 0.5, 1}") {
		t.Fatalf("expected custom buckets literal in output, got:\n%s", out)
	}
	if strings.Contains(out, "prometheus.DefBuckets") {
		t.Fatalf("expected DefBuckets to be absent when explicit buckets set")
	}
}
