//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what tracingWrapCalls 단위 테스트 (tracing.enabled AND wrap_calls 둘 다일 때만 true)
package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func mkFS(tr *manifest.ObservabilityTracing) *yongol.Fullstack {
	mf := &manifest.ProjectConfig{}
	if tr != nil {
		mf.Backend.Observability = &manifest.Observability{Tracing: tr}
	}
	return &yongol.Fullstack{Manifest: mf}
}
