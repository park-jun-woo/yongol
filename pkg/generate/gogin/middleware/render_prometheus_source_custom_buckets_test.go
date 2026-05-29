//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=observability
//ff:what TestRenderPrometheusSources_CustomBuckets — custom buckets 리터럴이 출력에 포함

package middleware

import (
	"strings"
	"testing"
)

func TestRenderPrometheusSources_CustomBuckets(t *testing.T) {
	files := renderPrometheusSources([]float64{0.1, 0.5, 1})
	combined := ""
	for _, v := range files {
		combined += v
	}
	if !strings.Contains(combined, "[]float64{0.1, 0.5, 1}") {
		t.Fatalf("expected custom buckets literal in output, got:\n%s", combined)
	}
	if strings.Contains(combined, "prometheus.DefBuckets") {
		t.Fatalf("expected DefBuckets to be absent when explicit buckets set")
	}
}
