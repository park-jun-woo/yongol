//ff:func feature=gen-ir type=test control=sequence
//ff:what modulePath/projectID/rateLimitHasEntries/corsIsEnabled/otelEnabled/prometheusEnabled/securityHeadersEnabled/hasAuthSequence manifest 추출·판정 헬퍼
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestOtelEnabled(t *testing.T) {
	if otelEnabled(nil) {
		t.Errorf("nil should be false")
	}
	off := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if otelEnabled(off) {
		t.Errorf("no observability should be false")
	}
	on := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		Observability: &manifest.Observability{Tracing: &manifest.ObservabilityTracing{Enabled: true}},
	}}}
	if !otelEnabled(on) {
		t.Errorf("tracing enabled should be true")
	}
}
