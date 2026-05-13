//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=observability
//ff:what TestRenderPrometheusSources_DefaultBuckets — 기본 buckets=nil → prometheus.DefBuckets

package middleware

import (
	"strings"
	"testing"
)

func TestRenderPrometheusSources_DefaultBuckets(t *testing.T) {
	files := renderPrometheusSources(nil)
	combined := ""
	for _, v := range files {
		combined += v
	}
	if !strings.Contains(combined, "prometheus.DefBuckets") {
		t.Fatalf("expected prometheus.DefBuckets literal for nil buckets, got:\n%s", combined)
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
		if !strings.Contains(combined, must) {
			t.Errorf("rendered source missing fragment %q", must)
		}
	}
}
