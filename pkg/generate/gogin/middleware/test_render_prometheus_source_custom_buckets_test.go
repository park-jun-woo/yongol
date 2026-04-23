//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestRenderPrometheusSource_CustomBuckets — custom buckets 리터럴이 출력에 포함

package middleware

import (
	"strings"
	"testing"
)

func TestRenderPrometheusSource_CustomBuckets(t *testing.T) {
	out := renderPrometheusSource([]float64{0.1, 0.5, 1})
	if !strings.Contains(out, "[]float64{0.1, 0.5, 1}") {
		t.Fatalf("expected custom buckets literal in output, got:\n%s", out)
	}
	if strings.Contains(out, "prometheus.DefBuckets") {
		t.Fatalf("expected DefBuckets to be absent when explicit buckets set")
	}
}
