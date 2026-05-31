//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what hasPrometheus — manifest.backend.observability.metrics.enabled (기본 true) 여부
package boot

import (
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithMetrics(m *pmanifest.ObservabilityMetrics) *yongol.Fullstack {
	return &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Observability: &pmanifest.Observability{Metrics: m}},
	}}
}
