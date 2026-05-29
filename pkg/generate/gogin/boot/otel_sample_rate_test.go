//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what otelSampleRate — tracing.sample_rate 값 결정 (0 이하는 1.0 기본값)

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsTracing(tr *pmanifest.ObservabilityTracing) *yongol.Fullstack {
	return &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Observability: &pmanifest.Observability{Tracing: tr}},
	}}
}

func TestOtelSampleRate(t *testing.T) {
	if got := otelSampleRate(nil); got != 1.0 {
		t.Errorf("no otel = %v, want 1.0", got)
	}
	// Enabled but zero rate → default 1.0 (explicit-0 semantics deferred to env).
	zero := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true, SampleRate: 0})
	if got := otelSampleRate(zero); got != 1.0 {
		t.Errorf("zero rate = %v, want 1.0", got)
	}
	neg := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true, SampleRate: -0.5})
	if got := otelSampleRate(neg); got != 1.0 {
		t.Errorf("negative rate = %v, want 1.0", got)
	}
	half := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true, SampleRate: 0.25})
	if got := otelSampleRate(half); got != 0.25 {
		t.Errorf("explicit rate = %v, want 0.25", got)
	}
}
