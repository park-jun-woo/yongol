//ff:func feature=validate type=test control=sequence topic=manifest-observability
//ff:what OBS-001 테스트 — metrics.path 누락 slash negative

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestObs01MetricsPath_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Metrics: &pmanifest.ObservabilityMetrics{Path: "metrics"},
				},
			},
		},
	}
	got := obs01MetricsPath(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
	if !strings.Contains(got[0].Message, "[OBS-001]") {
		t.Fatalf("message missing [OBS-001] prefix: %q", got[0].Message)
	}
}
