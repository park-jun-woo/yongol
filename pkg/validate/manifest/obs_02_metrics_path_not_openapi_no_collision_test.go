//ff:func feature=validate type=test control=sequence topic=manifest-observability
//ff:what OBS-002 테스트 — 서로 다른 path 는 충돌 없음

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestObs02MetricsPathNotOpenAPI_NoCollision(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: buildDocWithPath("/users"),
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Metrics: &pmanifest.ObservabilityMetrics{Path: "/metrics"},
				},
			},
		},
	}
	if got := obs02MetricsPathNotOpenAPI(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
