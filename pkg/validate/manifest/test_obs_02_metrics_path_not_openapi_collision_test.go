//ff:func feature=validate type=test control=sequence topic=manifest-observability
//ff:what OBS-002 테스트 — metrics.path 와 OpenAPI path 동일하면 ERROR

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestObs02MetricsPathNotOpenAPI_Collision(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: buildDocWithPath("/metrics"),
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Metrics: &pmanifest.ObservabilityMetrics{Path: "/metrics"},
				},
			},
		},
	}
	got := obs02MetricsPathNotOpenAPI(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
	if !strings.Contains(got[0].Message, "[OBS-002]") {
		t.Fatalf("message missing [OBS-002] prefix: %q", got[0].Message)
	}
}
