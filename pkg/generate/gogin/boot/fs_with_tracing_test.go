//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what hasOtel — manifest.backend.observability.tracing.enabled 여부 (기본 false)
package boot

import (
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithTracing(tr *pmanifest.ObservabilityTracing) *yongol.Fullstack {
	return &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Observability: &pmanifest.Observability{Tracing: tr}},
	}}
}
