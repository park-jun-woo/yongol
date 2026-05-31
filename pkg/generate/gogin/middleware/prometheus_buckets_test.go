//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증
package middleware

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestPrometheusBuckets(t *testing.T) {
	if got := prometheusBuckets(nil); got != nil {
		t.Errorf("nil fs: want nil, got %v", got)
	}
	if got := prometheusBuckets(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}); got != nil {
		t.Errorf("no observability: want nil, got %v", got)
	}
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Observability: &pmanifest.Observability{
			Metrics: &pmanifest.ObservabilityMetrics{Buckets: []float64{0.1, 1}},
		}},
	}}
	if got := prometheusBuckets(fs); len(got) != 2 || got[0] != 0.1 {
		t.Errorf("want [0.1 1], got %v", got)
	}
}
