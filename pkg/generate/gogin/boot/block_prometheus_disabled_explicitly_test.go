//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBlockPrometheus_DisabledExplicitly — metrics.enabled=false 시 inert block

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockPrometheus_DisabledExplicitly(t *testing.T) {
	disabled := false
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Observability: &pmanifest.Observability{
					Metrics: &pmanifest.ObservabilityMetrics{Enabled: &disabled},
				},
			},
		},
	}
	block := blockPrometheus(fs, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Fatalf("expected inert block when metrics disabled, got lines: %+v", block.Lines)
	}
}
