//ff:func feature=validate type=test-helper control=sequence topic=manifest-observability
//ff:what assertObs03ExporterAllowed — exporter 값 허용 여부 단언 헬퍼

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// assertObs03ExporterAllowed builds a fullstack with the given exporter value
// and asserts that obs03TracingExporter emits no diagnostic. Extracted to keep
// the iteration body small (Q4).
func assertObs03ExporterAllowed(t *testing.T, v string) {
	t.Helper()
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Tracing: &pmanifest.ObservabilityTracing{Enabled: true, Exporter: v},
				},
			},
		},
	}
	if got := obs03TracingExporter(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
