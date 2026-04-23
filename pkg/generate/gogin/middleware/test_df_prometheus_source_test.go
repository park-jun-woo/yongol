//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=observability
//ff:what TestRenderPrometheusSource_DefaultBuckets — 기본 buckets=nil → prometheus.DefBuckets

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
