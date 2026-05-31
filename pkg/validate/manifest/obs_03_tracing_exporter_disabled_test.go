//ff:func feature=validate type=test control=sequence topic=manifest-observability
//ff:what OBS-003 테스트 — tracing.enabled=false 면 exporter 값 미검증

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestObs03TracingExporter_Disabled(t *testing.T) {
	// Bogus value is tolerated while tracing.enabled is false — the
	// generator never reads exporter in that case.
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Tracing: &pmanifest.ObservabilityTracing{Enabled: false, Exporter: "bogus"},
				},
			},
		},
	}
	if got := obs03TracingExporter(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics when tracing disabled, got: %+v", got)
	}
}
