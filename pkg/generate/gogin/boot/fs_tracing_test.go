//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what otelSampleRate — tracing.sample_rate 값 결정 (0 이하는 1.0 기본값)
package boot

import (
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsTracing(tr *pmanifest.ObservabilityTracing) *yongol.Fullstack {
	return &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Observability: &pmanifest.Observability{Tracing: tr}},
	}}
}
