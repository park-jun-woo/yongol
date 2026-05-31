//ff:func feature=gen-ir type=test control=sequence
//ff:what modulePath/projectID/rateLimitHasEntries/corsIsEnabled/otelEnabled/prometheusEnabled/securityHeadersEnabled/hasAuthSequence manifest 추출·판정 헬퍼
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestPrometheusEnabled(t *testing.T) {
	// defaults to true (opt-out)
	if !prometheusEnabled(nil) {
		t.Errorf("nil should default true")
	}
	if !prometheusEnabled(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}) {
		t.Errorf("no observability should default true")
	}
	disabled := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		Observability: &manifest.Observability{Metrics: &manifest.ObservabilityMetrics{Enabled: boolPtr(false)}},
	}}}
	if prometheusEnabled(disabled) {
		t.Errorf("explicit false should be false")
	}
}
