//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what blockPrometheus — middleware.PrometheusMiddleware + /metrics 라우팅 등록
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockPrometheus_DisabledInert(t *testing.T) {
	disabled := false
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Observability: &pmanifest.Observability{
			Metrics: &pmanifest.ObservabilityMetrics{Enabled: &disabled},
		}},
	}}
	block := blockPrometheus(fs, "example.com/zenflow")
	if len(block.Lines) != 0 || len(block.Imports) != 0 {
		t.Fatalf("disabled prometheus must be inert, got %+v", block)
	}
}
