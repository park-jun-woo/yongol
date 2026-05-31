//ff:func feature=validate type=test control=sequence topic=manifest-observability
//ff:what OBS-004 테스트 — tracing.enabled=false 면 sample_rate 값 미검증

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestObs04TracingSampleRate_Disabled(t *testing.T) {
	// Bad value tolerated while tracing.enabled=false — no runtime impact.
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Tracing: &pmanifest.ObservabilityTracing{Enabled: false, SampleRate: 999},
				},
			},
		},
	}
	if got := obs04TracingSampleRate(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics when tracing disabled, got %+v", got)
	}
}
