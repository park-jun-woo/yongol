//ff:func feature=validate type=test control=sequence topic=manifest-observability
//ff:what OBS-002 테스트 — path 생략 시 기본 /metrics 도 충돌 판정

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestObs02MetricsPathNotOpenAPI_DefaultCollision(t *testing.T) {
	// path unset falls back to "/metrics" — still collides.
	fs := &yongol.Fullstack{
		OpenAPIDoc: buildDocWithPath("/metrics"),
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{Metrics: &pmanifest.ObservabilityMetrics{}},
			},
		},
	}
	if got := obs02MetricsPathNotOpenAPI(fs); len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
}
