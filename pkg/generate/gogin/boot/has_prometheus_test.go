//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what hasPrometheus — manifest.backend.observability.metrics.enabled (기본 true) 여부

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithMetrics(m *pmanifest.ObservabilityMetrics) *yongol.Fullstack {
	return &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Observability: &pmanifest.Observability{Metrics: m}},
	}}
}

func TestHasPrometheus(t *testing.T) {
	tru := true
	fals := false
	cases := []struct {
		name string
		fs   *yongol.Fullstack
		want bool
	}{
		{"nil fs defaults true", nil, true},
		{"nil manifest defaults true", &yongol.Fullstack{}, true},
		{"no metrics block defaults true", &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}, true},
		{"metrics nil enabled defaults true", fsWithMetrics(&pmanifest.ObservabilityMetrics{Enabled: nil}), true},
		{"explicitly enabled", fsWithMetrics(&pmanifest.ObservabilityMetrics{Enabled: &tru}), true},
		{"explicitly disabled", fsWithMetrics(&pmanifest.ObservabilityMetrics{Enabled: &fals}), false},
	}
	for _, c := range cases {
		if got := hasPrometheus(c.fs); got != c.want {
			t.Errorf("%s: hasPrometheus = %v, want %v", c.name, got, c.want)
		}
	}
}
