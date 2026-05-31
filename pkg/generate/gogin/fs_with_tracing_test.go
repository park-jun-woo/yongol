//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestResolveGoModDeps — tracing off/on(otlp/stdout/default) 의존성 병합 분기 검증
package gogin

import (
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithTracing(tr *pmanifest.ObservabilityTracing) *yongol.Fullstack {
	return &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{Tracing: tr},
			},
		},
	}
}
